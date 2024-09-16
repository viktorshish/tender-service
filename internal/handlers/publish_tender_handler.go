package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"tender-service/internal/config"
	"tender-service/internal/models"
)

func PublishTenderHandler(c *gin.Context) {
	// Получаем ID тендера из параметров URL
	tenderID := c.Param("id")

	// Преобразуем tenderID в UUID
	tenderUUID, err := uuid.Parse(tenderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID тендера"})
		return
	}

	// Ищем тендер по ID
	var tender models.Tender
	if err := config.DB.First(&tender, "id = ?", tenderUUID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тендер не найден"})
		return
	}

	// Получаем данные из запроса (кто пытается опубликовать тендер)
	var input struct {
		CreatorUsername string `json:"creatorUsername" binding:"required"`
	}

	// Привязываем данные из JSON-запроса
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации данных запроса"})
		return
	}

	// Ищем пользователя, который пытается опубликовать тендер
	var creator models.Employee
	if err := config.DB.Where("username = ?", input.CreatorUsername).First(&creator).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Проверяем, является ли пользователь доверенным лицом компании тендера
	var orgResp models.OrganizationResponsible
	if err := config.DB.Where("organization_id = ? AND user_id = ?", tender.OrganizationID, creator.ID).First(&orgResp).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Пользователь не является доверенным лицом организации"})
		return
	}

	// Меняем статус тендера на PUBLISHED
	tender.Status = models.StatusPublished

	// Сохраняем изменения в базе данных
	if err := config.DB.Save(&tender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса тендера"})
		return
	}

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
