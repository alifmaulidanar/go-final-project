package main

import (
	"fmt"
	"log"
	"net/http"

	"final-project/config"
	"final-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer db.Close()

	// Initialize the Gin router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		port := 8080 // Define the port
		message := fmt.Sprintf("Server running on %d", port)
		c.String(http.StatusOK, message)
	})

	// Setup routes
	routes.AdminRoutes(r, db)
	routes.ProductRoutes(r, db)
	routes.VariantRoutes(r, db)

	// Run the Gin server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
