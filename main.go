package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello World")

	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := route.Run(":8081")
	if err != nil {
		return
	}
}
