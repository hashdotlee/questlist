package models 

import (
	"github.com/jinzhu/gorm"
	)	

type Question struct {
	gorm.Model
	Title string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID uint `json:"user_id"`
	Upvote int `gorm:"not null" json:"upvote"`
	Downvote int `gorm:"not null" json:"downvote"`
	Answers []*Answer `gorm:"foreignkey:QuestionID" json:"answers"`
	Topics []*Topic `gorm:"many2many:topic_question" json:"topics"`
}
