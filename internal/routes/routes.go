package routes

import (
	"github.com/gin-gonic/gin"

	"tender-service/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.GET("/ping", handlers.PingHandler)

		api.POST("/tenders/new", handlers.CreateTenderHandler)

		api.PATCH("/tenders/:id/publish", handlers.PublishTenderHandler)

		api.GET("/tenders", handlers.GetTendersHandler)

		api.PATCH("/tenders/:id/cancel", handlers.CancelTenderHandler)
	}
}
