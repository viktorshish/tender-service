package routes

import (
	"github.com/gin-gonic/gin"
	"tender-service/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	apiTenders := router.Group("/api/tenders")
	{
		apiTenders.POST("/new", handlers.CreateTender) // Responsible for the organization

		apiTenders.PUT("/:id/publish", handlers.PublishTender) // Responsible for the organization

		apiTenders.GET("/", handlers.GetTenders) // All users

		//apiTenders.GET("/my", handlers.GetMyTenders) // All users

		//apiTenders.GET("/service", handlers.GetTendersByService) // All users

		//apiTenders.PUT("/:id/cancel", handlers.CancelTender) // Responsible for the organization

		//apiTenders.PUT("/:id/edit", handlers.EditTender) // Responsible for the organization

		//apiTenders.PUT("/:id/rollback", handlers.RollbackTender) // Responsible for the organization")
	}

}
