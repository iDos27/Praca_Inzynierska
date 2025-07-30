package models

import "time"

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Emoji       string    `json:"emoji"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type MenuItem struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
}

type Order struct {
	ID          int         `json:"id"`
	TableNumber string      `json:"table_number"`
	TotalAmount float64     `json:"total_amount"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Items       []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
	ID           int     `json:"id"`
	OrderID      int     `json:"order_id"`
	MenuItemID   int     `json:"menu_item_id"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
	TotalPrice   float64 `json:"total_price"`
	MenuItemName string  `json:"menu_item_name,omitempty"`
}

type CreateOrderRequest struct {
	TableNumber string `json:"table_number"`
	Items       []struct {
		ID       int `json:"id"`
		Quantity int `json:"quantity"`
	} `json:"items"`
}
