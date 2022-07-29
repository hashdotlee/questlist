package models 

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `gorm:"unique" json:"email"`
	Role string ` json:"role"`
	Phone string `json:"phone"`
	Birthday string `json:"birthday"`
	Verified bool `json:"verified"`
	Address string `json:"address"`
}
