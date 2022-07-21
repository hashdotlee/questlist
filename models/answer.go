package models

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	QuestionID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Content string `gorm:"not null"`
	IsCorrect bool `gorm:"not null"`
	Upvote int `gorm:"not null"`
}
