package handlers

import (
	"fmt"
	"net/http"
	"restaurant-backend/database"
	"restaurant-backend/models"

	"github.com/gin-gonic/gin"
)

// Pobieranie kategorie
func GetCategories(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, emoji, description, created_at FROM categories ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania kategorii"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Emoji, &category.Description, &category.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd skanowania kategorii"})
			return
		}
		categories = append(categories, category)
	}
	c.JSON(http.StatusOK, categories)
}

// Pobieranie menu dla kategorii
func GetMenuItems(c *gin.Context) {
	categoryID := c.Param("categoryId")
	rows, err := database.DB.Query(`
		SELECT id, category_id, name, description, price, is_available, created_at
		FROM menu_items
		WHERE category_id = $1 AND is_available = true
		ORDER BY id`, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania menu"})
		return
	}
	defer rows.Close()

	var menuItems []models.MenuItem
	for rows.Next() {
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.CategoryID, &item.Name, &item.Description, &item.Price, &item.IsAvailable, &item.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd skanowania menu"})
			return
		}
		menuItems = append(menuItems, item)
	}
	c.JSON(http.StatusOK, menuItems)
}

// Pobieranie całego menu
func GetAllMenuItems(c *gin.Context) {
	rows, err := database.DB.Query(`
        SELECT m.id, m.category_id, m.name, m.description, m.price, m.is_available, m.created_at 
        FROM menu_items m 
        WHERE m.is_available = true 
        ORDER BY m.category_id, m.id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania menu"})
		return
	}
	defer rows.Close()

	menuByCategory := make(map[int][]models.MenuItem)
	for rows.Next() {
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.CategoryID, &item.Name, &item.Description, &item.Price, &item.IsAvailable, &item.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd skanowania pozycji menu"})
			return
		}
		menuByCategory[item.CategoryID] = append(menuByCategory[item.CategoryID], item)
	}

	c.JSON(http.StatusOK, menuByCategory)
}

func CreateOrder(c *gin.Context) {
	fmt.Printf("DEBUG: CreateOrder została wywołana\n")

	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("DEBUG: Błąd parsowania JSON: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nieprawidłowe dane"})
		return
	}

	fmt.Printf("DEBUG: Otrzymane dane: %+v\n", req)

	if req.OrderType != "table" && req.OrderType != "pickup" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowy typ zamówienia"})
		return
	}

	if req.OrderType == "table" && req.TableNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Numer stolika jest wymagany"})
		return
	}

	// Generuj numer zamówienia w przedziale 1-99 (cyklicznie) PRZED transakcją
	// Znajdź najnowszy rekord i jego order_number
	var orderNumber int
	err := database.DB.QueryRow("SELECT COALESCE(order_number, 0) FROM orders ORDER BY id DESC LIMIT 1").Scan(&orderNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// Tabela jest pusta, zacznij od 0 (będzie inkrementowane do 1)
			orderNumber = 0
			fmt.Printf("DEBUG: Tabela orders jest pusta, zaczynam od 0\n")
		} else {
			fmt.Printf("DEBUG: Błąd przy pobieraniu order_number: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd generowania numeru"})
			return
		}
	} else {
		fmt.Printf("DEBUG: Zapytanie przeszło pomyślnie\n")
	}

	// DEBUG: sprawdź co zwraca MAX i kilka ostatnich rekordów
	fmt.Printf("DEBUG: MAX order_number z bazy: %d\n", orderNumber)

	// Sprawdź ostatnie 3 rekordy
	rows, _ := database.DB.Query("SELECT id, order_number FROM orders ORDER BY id DESC LIMIT 3")
	defer rows.Close()
	fmt.Printf("DEBUG: Ostatnie 3 rekordy w bazie:\n")
	for rows.Next() {
		var id, num int
		rows.Scan(&id, &num)
		fmt.Printf("  ID: %d, order_number: %d\n", id, num)
	}

	// Następny numer w cyklu 0-99
	orderNumber++
	if orderNumber > 99 {
		orderNumber = 0
	}

	// DEBUG: sprawdź jaki numer będziemy zapisywać
	fmt.Printf("DEBUG: Nowy order_number do zapisania: %d\n", orderNumber)

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd rozpoczynania transakcji"})
		return
	}
	defer tx.Rollback()

	var orderID int
	var totalAmount float64

	for _, item := range req.Items {
		var price float64
		err := tx.QueryRow("SELECT price FROM menu_items WHERE id = $1", item.ID).Scan(&price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Nieprawidłowe ID"})
			return
		}
		totalAmount += price * float64(item.Quantity)
	}

	err = tx.QueryRow(`
		INSERT INTO orders (order_type, table_number, order_number, total_amount, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, 'pending', NOW(), NOW())
		RETURNING id`,
		req.OrderType, req.TableNumber, orderNumber, totalAmount).Scan(&orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd tworzenia zamówienia"})
		return
	}

	for _, item := range req.Items {
		var price float64
		err := tx.QueryRow("SELECT price FROM menu_items WHERE id = $1", item.ID).Scan(&price)
		if err != nil {
			continue
		}

		totalPrice := price * float64(item.Quantity)
		_, err = tx.Exec(`
            INSERT INTO order_items (order_id, menu_item_id, quantity, unit_price, total_price) 
            VALUES ($1, $2, $3, $4, $5)`,
			orderID, item.ID, item.Quantity, price, totalPrice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd dodawania pozycji zamówienia"})
			return
		}
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd zatwierdzania zamówienia"})
		return
	}

	// DEBUG: Pokaż utworzone zamówienie
	fmt.Printf("DEBUG: Utworzono zamówienie - ID: %d, Numer: %d\n", orderID, orderNumber)

	response := gin.H{
		"message":      "Zamówienie zostało utworzone",
		"order_id":     orderID,
		"total":        totalAmount,
		"order_number": orderNumber,
	}

	c.JSON(http.StatusCreated, response)
}

func GetOrders(c *gin.Context) {
	rows, err := database.DB.Query(`
        SELECT id, order_type, table_number, order_number, total_amount, status, created_at, updated_at 
        FROM orders 
        ORDER BY created_at DESC`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania zamówień"})
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.OrderType, &order.TableNumber, &order.OrderNumber, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd skanowania zamówień"})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}
