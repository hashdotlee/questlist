package models

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	QuestionID uint `gorm:"not null" json:"question_id"`
	UserID uint `json:"user_id"`
	Content string `gorm:"not null" json:"content"`
	IsCorrect bool `gorm:"not null" json:"is_correct"`
	Upvote int `gorm:"not null" json:"upvote"`
	Downvote int `gorm:"not null" json:"downvote"`
	Verified bool `json:"verified"`
}
