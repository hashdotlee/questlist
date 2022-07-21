package models 

import (
	"github.com/jinzhu/gorm"
	)	

type Question struct {
	gorm.Model
	Title string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Upvote int `gorm:"not null"`
	Downvote int `gorm:"not null"`
	Answers []Answer `gorm:"foreignkey:QuestionID"`
	Topics []*Topic `gorm:"many2many:topic_question"`
}
