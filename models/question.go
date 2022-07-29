package models 

import (
	"github.com/jinzhu/gorm"
	)	

	type QuestionType int 
	const (
		QuestionTypePublic QuestionType = iota
		QuestionTypePrivate
	)

type Question struct {
	gorm.Model
	Title string `gorm:"not null" json:"title"`
	Image string `json:"image"`
	Content string `gorm:"not null" json:"content"`
	UserID uint `json:"user_id"`
	Type QuestionType `json:"type"`
	Vote []VoteQuestion `json:"vote"`
	Answers []*Answer `gorm:"foreignkey:QuestionID" json:"answers"`
	Topics []Topic `gorm:"many2many:topic_question" json:"topics"`
}
