package main

import (
	"ecommerce/api-gateway/config"
	"ecommerce/api-gateway/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	clients, err := config.NewClients()
	if err != nil {
		log.Fatalf("Failed to initialize clients: %v", err)
	}
	defer clients.Close()

	r := gin.Default()

	metrics := handlers.NewMetrics()
	r.Use(handlers.Logger(), handlers.Auth(), metrics.Track())

	h := handlers.NewHandler(clients)

	// Inventory routes
	r.POST("/products", h.CreateProduct)
	r.GET("/products/:id", h.GetProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	r.GET("/products", h.ListProducts)

	// Order routes
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders/:id", h.GetOrder)
	r.PUT("/orders/:id/status", h.UpdateOrderStatus)
	r.GET("/users/:id/orders", h.ListUserOrders)

	// Metrics endpoint
	r.GET("/metrics", metrics.GetMetrics)

	log.Println("API Gateway running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
