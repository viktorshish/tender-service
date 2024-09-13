package models

import "time"

type OrganizationType string

const (
	OrganizationTypeIE  OrganizationType = "IE"
	OrganizationTypeLLC OrganizationType = "LLC"
	OrganizationTypeJSC OrganizationType = "JSC"
)

type Organization struct {
	ID          string           `gorm:"type:uuid;primary_key" json:"id"`
	Name        string           `gorm:"type:varchar(100);not null" json:"name"`
	Description string           `gorm:"type:text" json:"description"`
	Type        OrganizationType `gorm:"type:organization_type" json:"type"`
	CreatedAt   time.Time        `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (Organization) TableName() string {
	return "organization"
}
