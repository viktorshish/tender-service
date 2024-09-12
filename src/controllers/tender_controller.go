package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tender_service_project/src/services"
)

func CreateTenderHandler(c *gin.Context) {
	var input services.CreateTenderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tender, err := services.CreateTender(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}
