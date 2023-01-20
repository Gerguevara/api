package models

import "gorm.io/gorm"

type User struct {
	// lee el struct y lo convierte en una tabla
	gorm.Model

	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null" `
	Email     string `json:"email" gorm:"not null;uniqueIndex"`
	Tasks     []Task `json:"tasks"`
}
