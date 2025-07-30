package main

import (
	"log"
	"restaurant-backend/database"
	"restaurant-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	defer database.CloseDB()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // React klient
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		// Menu endpoints
		api.GET("/categories", handlers.GetCategories)
		api.GET("/menu/:categoryId", handlers.GetMenuItems)
		api.GET("/menu", handlers.GetAllMenuItems)

		// Order endpoints
		api.POST("/orders", handlers.CreateOrder)
		api.GET("/orders", handlers.GetOrders)
	}

	log.Println("Serwer działa na http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Błąd uruchamiania serwera:", err)
	}

}
