package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/assam5649/Siosabin-go/pkg/config"
)

func main() {

	err := config.wait()

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

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
