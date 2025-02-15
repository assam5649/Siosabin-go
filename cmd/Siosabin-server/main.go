package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Go Server!",
		})
	})

	port := "5000"
	fmt.Println("Starting Gin server on port ", port)

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
