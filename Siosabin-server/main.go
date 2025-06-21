package main

import (
	"fmt"
	"log"

	"Siosabin-go/Siosabin-server/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Wait()

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
