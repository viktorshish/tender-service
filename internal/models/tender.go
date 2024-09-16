package models

import (
	"time"

	"github.com/google/uuid"
)

type StatusType string

const (
	StatusCreated   StatusType = "CREATED"
	StatusPublished StatusType = "PUBLISHED"
	StatusClosed    StatusType = "CLOSED"
	StatusCanceled  StatusType = "CANCELED"
)

type Tender struct {
	ID             uuid.UUID               `gorm:"type:uuid;primary_key" json:"id"`
	Name           string                  `gorm:"type:varchar(255);not null" json:"name"`
	Description    string                  `gorm:"type:text" json:"description"`
	ServiceType    string                  `gorm:"type:varchar(255);not null" json:"serviceType"`
	Status         StatusType              `gorm:"type:status_type;not null" json:"status"`
	Version        int                     `gorm:"default:1" json:"version"`
	OrganizationID uuid.UUID               `gorm:"type:uuid;not null" json:"organizationId"`
	Organization   Organization            `gorm:"foreignKey:OrganizationID" json:"organization"`
	ResponsibleID  string                  `gorm:"type:uuid;not null" json:"responsibleId"`
	Responsible    OrganizationResponsible `gorm:"foreignKey:ResponsibleID" json:"responsible"`
	CreatorID      string                  `gorm:"type:uuid;not null" json:"creatorId"`
	Creator        Employee                `gorm:"foreignKey:CreatorID" json:"creator"`
	CreatedAt      time.Time               `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time               `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Tender) TableName() string {
	return "tender"
}
