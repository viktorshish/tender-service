package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(router.Run(":8080"))
}
