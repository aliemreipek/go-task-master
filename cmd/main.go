package main

import (
	"log"

	"github.com/aliemreipek/go-task-master/pkg/config"
	"github.com/aliemreipek/go-task-master/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize Database Connection
	config.Connect()

	// 2. Initialize Gin Router
	r := gin.Default()

	// 3. Setup Routes
	routes.SetupRoutes(r)

	// 4. Start Server
	log.Println("Server is running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
