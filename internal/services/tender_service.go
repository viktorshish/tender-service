package services

import (
	"errors"
	"tender-service/internal/config"
	"tender-service/internal/models"
)

type CreateTenderInput struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description"`
	ServiceType     string `json:"serviceType" binding:"required"`
	Status          string `json:"status" binding:"required"`
	OrganizationID  string `json:"organizationId" binding:"required,uuid"`
	CreatorUsername string `json:"creatorUsername" binding:"required"`
}

func CreateTender(input CreateTenderInput) (*models.Tender, error) {
	var employee models.Employee

	if err := config.DB.Where("username = ?", input.CreatorUsername).First(&employee).Error; err != nil {
		return nil, errors.New("user not found")
	}

	var orgResp models.OrganizationResponsible
	if err := config.DB.Where("user_id = ? AND organization_id = ?", employee.ID, input.OrganizationID).First(&orgResp).Error; err != nil {
		return nil, errors.New("user is not responsible for this organization")
	}

	tender := models.Tender{
		Name:           input.Name,
		Description:    input.Description,
		ServiceType:    input.ServiceType,
		Status:         models.StatusType(input.Status),
		Version:        1,
		OrganizationID: input.OrganizationID,
		ResponsibleID:  employee.ID,
		CreatorID:      employee.ID,
	}

	if err := config.DB.Create(&tender).Error; err != nil {
		return nil, errors.New("failed to create tender")
	}

	return &tender, nil
}
