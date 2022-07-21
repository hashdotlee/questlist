package models 

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email string `gorm:"not null"`
	Role string `gorm:"not null"`
	Birthday string `gorm:"not null"`
	Verified bool `gorm:"not null"`
}
