package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"tender-service/internal/config"
	"tender-service/internal/models"
)

func PublishTender(c *gin.Context) {
	tenderID := c.Param("id")

	tenderUUID, err := uuid.Parse(tenderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID тендера"})
		return
	}

	var tender models.Tender
	if err := config.DB.First(&tender, "id = ?", tenderUUID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тендер не найден"})
		return
	}

	var input struct {
		CreatorUsername string `json:"creatorUsername" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации данных запроса"})
		return
	}

	var creator models.Employee
	if err := config.DB.Where("username = ?", input.CreatorUsername).First(&creator).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	var orgResp models.OrganizationResponsible
	if err := config.DB.Where("organization_id = ? AND user_id = ?", tender.OrganizationID, creator.ID).First(&orgResp).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Пользователь не является доверенным лицом организации"})
		return
	}

	tender.Status = models.StatusPublished

	if err := config.DB.Save(&tender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса тендера"})
		return
	}

	log.Println("Tender published")

	responseDTO := TenderDTO{
		ID:             tender.ID,
		Name:           tender.Name,
		Description:    tender.Description,
		ServiceType:    tender.ServiceType,
		OrganizationID: tender.OrganizationID,
		CreatorID:      tender.CreatorID,
		Status:         tender.Status,
	}

	c.JSON(http.StatusOK, responseDTO)
}
