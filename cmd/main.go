package main

import (
	"log"
	"gin-sqlite-gorm-test/internal/database"
	"gin-sqlite-gorm-test/internal/routes"
)

func main() {
	// Initialize database
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Set up router
	r := routes.SetupRouter()

	// Run the server
	r.Run(":8080")
}