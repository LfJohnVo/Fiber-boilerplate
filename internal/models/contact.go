package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Enterprise *string `json:"empresa" gorm:"default:null"` // Opcional
	Message    string  `json:"mensaje"`
	Platform   string  `json:"platform"`
}
