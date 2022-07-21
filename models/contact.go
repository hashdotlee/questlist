package models 

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Info string `gorm:"not null" json:"info"`
	Type string `gorm:"not null" json:"type"`
	UserID uint `gorm:"not null" json:"user_id"`
}
