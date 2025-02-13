package main

import (
	database "Text-Search-Engine/Database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up the Gin router
	r := gin.Default()

	// Set up the routes
	r.GET("/search", database.SearchBooks)

	// Run the server
	r.Run(":8080")
}
