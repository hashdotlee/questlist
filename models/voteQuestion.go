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
	QuestionID uint `gorm:"primaryKey" json:"answer_id"`
	UserID uint `json:"user_id" gorm:"primaryKey"`
	Type VoteQuestionType `json:"type"`
}

