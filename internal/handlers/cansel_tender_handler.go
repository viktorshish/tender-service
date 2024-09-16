package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"tender-service/internal/config"
	"tender-service/internal/models"
)

func CancelTenderHandler(c *gin.Context) {
	// Получаем ID тендера из параметров URL
	tenderID := c.Param("id")
	tenderUUID, err := uuid.Parse(tenderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID тендера"})
		return
	}

	// Получаем данные пользователя, который пытается отменить тендер
	var input struct {
		CreatorUsername string `json:"creatorUsername" binding:"required"`
	}

	// Привязываем данные из JSON-запроса
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации данных запроса"})
		return
	}

	// Ищем тендер по ID
	var tender models.Tender
	if err := config.DB.Preload("Creator").Preload("Organization").Where("id = ?", tenderUUID).First(&tender).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тендер не найден"})
		return
	}

	// Ищем пользователя, который пытается отменить тендер
	var user models.Employee
	if err := config.DB.Where("username = ?", input.CreatorUsername).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Проверяем, является ли пользователь автором или доверенным лицом организации
	var orgResp models.OrganizationResponsible
	if tender.CreatorID != user.ID {
		// Проверяем, является ли пользователь ответственным за организацию
		if err := config.DB.Where("organization_id = ? AND user_id = ?", tender.OrganizationID, user.ID).First(&orgResp).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Пользователь не автор и не ответственное лицо за организацию"})
			return
		}
	}

	// Логируем процесс отмены тендера
	log.Printf("Отмена тендера %s от имени пользователя %s", tenderID, user.Username)

	tender.Status = models.StatusCanceled

	// Сохраняем изменения в базе данных
	if err := config.DB.Save(&tender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса тендера"})
		return
	}

	// Возвращаем обновлённые данные тендера
	c.JSON(http.StatusOK, gin.H{
		"message":  "Тендер отменён",
		"tenderId": tender.ID,
		"status":   tender.Status,
	})
}
