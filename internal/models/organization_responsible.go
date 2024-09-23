package models

import "github.com/google/uuid"

type OrganizationResponsible struct {
	ID             uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrganizationID string       `gorm:"type:uuid" json:"organization_id"`
	UserID         string       `gorm:"type:uuid" json:"user_id"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"organization"`
	User           Employee     `gorm:"foreignKey:UserID; references:ID;constraint:OnDelete:CASCADE" json:"user"`
}

func (OrganizationResponsible) TableName() string {
	return "organization_responsible"
}
