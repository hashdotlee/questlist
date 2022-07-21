package models 

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Info string `gorm:"not null"`
	Type string `gorm:"not null"`
	UserID uint `gorm:"not null"`
}
