package models

import (
	"time"
)

type Employee struct {
	ID        string    `gorm:"type:uuid;primary_key" json:"id"`
	Username  string    `gorm:"type:varchar(50);unique;not null" json:"username"`
	Firstname string    `gorm:"type:varchar(50)" json:"first_name"`
	Lastname  string    `gorm:"type:varchar(50)" json:"last_name"`
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

// Явное указание имени таблицы
func (Employee) TableName() string {
	return "employee"
}
