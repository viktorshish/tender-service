package routes

import (
	"github.com/gin-gonic/gin"

	"tender-service/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/tenders/new", handlers.CreateTenderHandler)

	router.POST("/api/tenders/:tenderId/publish", handlers.PublishTenderHandler)
}
