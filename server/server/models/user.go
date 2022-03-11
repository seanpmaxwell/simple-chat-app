package models

import (
	"gorm.io/gorm"
)

// User
type User struct {
	gorm.Model
	Email     string    `json:"email" gorm:"unique;size:255;not null"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	UserCreds UserCreds `gorm:"constraint:OnDelete:CASCADE;"`
}

// Store sensitive info info
type UserCreds struct {
	gorm.Model
	Pwdhash string `gorm:"size:255;not null"`
	UserID  uint
}
