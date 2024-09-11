package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"tender-service/src/config"
)

func main() {
	config.LoadConfig()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	serverAddress := os.Getenv("SERVER_ADDRESS")
	log.Fatal(router.Run(serverAddress))
}
