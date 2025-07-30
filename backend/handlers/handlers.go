package handlers

import (
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
	var request models.CreateOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nieprawidłowe żądanie"})
		return
	}
	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd Rozpoczynania transakcji"})
		return
	}
	defer tx.Rollback()

	var orderID int
	var totalAmount float64

	for _, item := range request.Items {
		var price float64
		err := tx.QueryRow("SELECT price FROM menu_items WHERE id = $1", item.ID).Scan(&price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Nieprawidłowe ID"})
			return
		}
		totalAmount += price * float64(item.Quantity)
	}

	err = tx.QueryRow(`
		INSERT INTO orders (table_number, total_amount, status, created_at, updated_at)
		VALUES ($1, $2, 'pending', NOW(), NOW())
		RETURNING id`,
		request.TableNumber, totalAmount).Scan(&orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd tworzenia zamówienia"})
		return
	}

	for _, item := range request.Items {
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

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Zamówienie zostało utworzone",
		"order_id": orderID,
		"total":    totalAmount,
	})
}

func GetOrders(c *gin.Context) {
	rows, err := database.DB.Query(`
        SELECT id, table_number, total_amount, status, created_at, updated_at 
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
		err := rows.Scan(&order.ID, &order.TableNumber, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd skanowania zamówień"})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}
