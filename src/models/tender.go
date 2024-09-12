package models

import "time"

type StatusType string

const (
	StatusCreated   StatusType = "CREATED"
	StatusPublished StatusType = "PUBLISHED"
	StatusClosed    StatusType = "CLOSED"
	StatusCanceled  StatusType = "CANCELED"
)

type Tender struct {
	ID             string                  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name           string                  `gorm:"type:varchar(255);not null; " json:"title"`
	Description    string                  `gorm:"type:text;" json:"description"`
	Status         StatusType              `gorm:"type:status_type" json:"status"`
	Version        int                     `gorm:"type:int; default:1" json:"version"`
	OrganizationID string                  `gorm:"type:uuid;not null" json:"organizationId"`
	Organization   Organization            `gorm:"foreignKey:OrganizationID" json:"organization"`
	ResponsibleID  string                  `gorm:"type:uuid;not null" json:"responsibleId"`
	Responsible    OrganizationResponsible `gorm:"foreignKey:ResponsibleID" json:"responsible"`
	CreatedAt      time.Time               `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time               `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}
