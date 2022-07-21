package models 

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email string `gorm:"not null" json:"email"`
	Role string `gorm:"not null" json:"role"`
	Birthday string `gorm:"not null" json:"birthday"`
	Verified bool `gorm:"not null" json:"verified"`
	Contacts []Contact `gorm:"foreignkey:UserID" json:"contacts"`
}
