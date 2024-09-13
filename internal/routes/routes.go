package routes

import (
	"github.com/gin-gonic/gin"

	"tender-service/internal/handlers"
)

// SetupRoutes настраивает маршруты для приложения.
func SetupRoutes(router *gin.Engine) {
	// Маршрут для создания тендера.
	router.POST("/api/tenders/new", handlers.CreateTenderHandler)
}
