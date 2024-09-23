package models

import (
	"github.com/google/uuid"
	"time"
)

type OrganizationType string

const (
	OrganizationTypeIE  OrganizationType = "IE"
	OrganizationTypeLLC OrganizationType = "LLC"
	OrganizationTypeJSC OrganizationType = "JSC"
)

type Organization struct {
	ID          uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string           `gorm:"type:varchar(100);not null" json:"name"`
	Description string           `gorm:"type:text" json:"description"`
	Type        OrganizationType `gorm:"type:organization_type" json:"type"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

func (Organization) TableName() string {
	return "organization"
}
