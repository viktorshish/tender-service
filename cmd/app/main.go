package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"tender-service/internal/config"
	"tender-service/internal/routes"
)

func main() {
	config.LoadConfig()

	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	/*
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})*/

	serverAddress := os.Getenv("SERVER_ADDRESS")
	log.Fatal(router.Run(serverAddress))
}
