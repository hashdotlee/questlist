package models

import (
	"github.com/jinzhu/gorm"
)

type VoteQuestionType int

const (
	UpVoteQuestion VoteQuestionType = iota
	DownVoteQuestion
)

type VoteQuestion struct {
	gorm.Model
	QuestionID uint `gorm:"not null" json:"answer_id"`
	UserID uint `json:"user_id" gorm:"not null"`
	Type VoteQuestionType `json:"type"`
}

