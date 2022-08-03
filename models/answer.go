package models

import (
	"github.com/jinzhu/gorm"
)

type Vote int 
const (
	UpVote Vote = iota
	DownVote
)

type Answer struct {
	gorm.Model
	QuestionID uint `gorm:"not null" json:"question_id"`
	UserID uint `json:"user_id"`
	Content string `gorm:"not null" json:"content"`
	IsCorrect bool `gorm:"not null" json:"is_correct"`
	Vote []VoteAnswer `json:"vote" gorm:"foreignkey:AnswerID"`
	Verified bool `json:"verified"`
}
