package models

import (
	"time"
)

type VoteQuestionType int

const (
	UpVoteQuestion VoteQuestionType = iota
	DownVoteQuestion
)

type VoteQuestion struct {
	ID uint `gorm:"autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"updated_at"`
	QuestionID uint `gorm:"primaryKey" json:"answer_id"`
	UserID uint `json:"user_id" gorm:"primaryKey"`
	Type VoteQuestionType `json:"type"`
}

