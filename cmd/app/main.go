package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"tender-service/internal/config"
	"tender-service/internal/routes"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	log.Fatal(router.Run(":8080"))
}
