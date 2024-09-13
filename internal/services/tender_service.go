package services

import (
	"errors"
	"github.com/google/uuid"
	"tender-service/internal/config"
	"tender-service/internal/models"
)

type CreateTenderInput struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ServiceType     string `json:"serviceType"`
	Status          string `json:"status"`
	OrganizationID  string `json:"organizationId"`
	CreatorUsername string `json:"creatorUsername"`
}

type TenderResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceType string `json:"serviceType"`
	Status      string `json:"status"`
}

func CreateTender(input CreateTenderInput) (*TenderResponse, error) {
	var employee models.Employee
	if err := config.DB.Where("username = ?", input.CreatorUsername).First(&employee).Error; err != nil {
		return nil, errors.New("user not found")
	}

	var orgResp models.OrganizationResponsible
	if err := config.DB.Where("user_id = ? AND organization_id = ?", employee.ID, input.OrganizationID).First(&orgResp).Error; err != nil {
		return nil, errors.New("user is not responsible for this organization")
	}

	tender := models.Tender{
		ID:             uuid.New(),
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

	response := &TenderResponse{
		ID:          tender.ID.String(),
		Name:        tender.Name,
		Description: tender.Description,
		ServiceType: tender.ServiceType,
		Status:      string(tender.Status),
	}

	return response, nil
}

func PublishTender(tenderID string, username string) (*TenderResponse, error) {
	var tender models.Tender
	var employee models.Employee

	if err := config.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		return nil, errors.New("user not found")
	}

	var orgResp models.OrganizationResponsible
	if err := config.DB.Where("user_id = ? AND organization_id = ?", employee.ID, tender.OrganizationID).First(&orgResp).Error; err != nil {
		return nil, errors.New("user is not responsible for this organization")
	}

	tender.Status = "PUBLISHED"

	if err := config.DB.Save(&tender).Error; err != nil {
		return nil, errors.New("failed to publish tender")
	}

	response := &TenderResponse{
		ID:          tender.ID.String(),
		Name:        tender.Name,
		Description: tender.Description,
		ServiceType: tender.ServiceType,
		Status:      string(tender.Status),
	}

	return response, nil
}
