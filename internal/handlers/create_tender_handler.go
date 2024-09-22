package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"tender-service/internal/config"
	"tender-service/internal/models"
)

type CreateTenderInput struct {
	Name            string    `json:"name" binding:"required"`
	Description     string    `json:"description"`
	ServiceType     string    `json:"serviceType"`
	OrganizationID  uuid.UUID `json:"organizationId" binding:"required"`
	CreatorUsername string    `json:"creatorUsername" binding:"required"`
}

type TenderDTO struct {
	ID             uuid.UUID         `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	ServiceType    string            `json:"serviceType"`
	OrganizationID uuid.UUID         `json:"organizationId"`
	CreatorID      string            `json:"creatorUsername"`
	Status         models.StatusType `json:"status"`
}

func CreateTender(c *gin.Context) {
	var input CreateTenderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var creator models.Employee
	if err := config.DB.Where("username = ?", input.CreatorUsername).First(&creator).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Проверка, что пользователь является ответственным за организацию
	var orgResp models.OrganizationResponsible
	if err := config.DB.Where("organization_id = ? AND user_id = ?", input.OrganizationID, creator.ID).First(&orgResp).Error; err != nil {
		log.Println("Пользователь не является ответственным за организацию:", creator.ID)
		c.JSON(http.StatusForbidden, gin.H{"error": "Пользователь не является ответственным за организацию"})
		return
	}

	tender := models.Tender{
		ID:             uuid.New(),
		Name:           input.Name,
		Description:    input.Description,
		ServiceType:    input.ServiceType,
		Status:         models.StatusCreated,
		Version:        1,
		OrganizationID: input.OrganizationID,
		ResponsibleID:  creator.ID,
		CreatorID:      creator.ID,
	}

	if err := config.DB.Create(&tender).Error; err != nil {
		log.Println("Ошибка при создании тендера:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Тендер создан успешно:", tender.ID)

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
