package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"tender-service/internal/config"
	"tender-service/internal/models"
)

type TenderResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ServiceType  string    `json:"serviceType"`
	Status       string    `json:"status"`
	Organization string    `json:"organizationName"`
}

func GetTendersHandler(c *gin.Context) {
	// Получаем параметр фильтрации serviceType (если он есть)
	serviceType := c.Query("serviceType")

	var tenders []models.Tender

	// Загружаем тендеры по фильтру
	if serviceType != "" {
		// Если передан параметр serviceType, фильтруем по типу услуг и загружаем связанные данные
		if err := config.DB.Preload("Organization").Preload("Creator").Where("service_type = ?", serviceType).Find(&tenders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении тендеров"})
			return
		}
	} else {
		// Если параметр не передан, загружаем все тендеры с связанными данными
		if err := config.DB.Preload("Organization").Preload("Creator").Find(&tenders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении тендеров"})
			return
		}
	}

	var tenderResponses []TenderResponse

	for _, tender := range tenders {
		tenderResponses = append(tenderResponses, TenderResponse{
			ID:           tender.ID,
			Name:         tender.Name,
			Description:  tender.Description,
			ServiceType:  tender.ServiceType,
			Status:       string(tender.Status),    // Преобразуем статус в строку
			Organization: tender.Organization.Name, // Имя организации
		})
	}
	c.JSON(http.StatusOK, tenderResponses)
}
