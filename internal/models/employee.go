package models

import (
	"github.com/google/uuid"
	"time"
)

type Employee struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Username  string    `gorm:"type:varchar(50);unique;not null" json:"username"`
	Firstname string    `gorm:"type:varchar(50)" json:"first_name"`
	Lastname  string    `gorm:"type:varchar(50)" json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Employee) TableName() string {
	return "employee"
}
