package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Nome  string `gorm:"not null" json:"nome"`
	Email string `gorm:"not null;unique" json:"email"`
}
