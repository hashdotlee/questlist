package models

import (
	"time"
)

type VoteAnswerType int

const (
	UpVoteAnswer VoteAnswerType = iota
	DownVoteAnswer
)

type VoteAnswer struct {
	ID uint `gorm:"autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"updated_at"`
	AnswerID uint `gorm:"primaryKey;autoIncrement:false" json:"answer_id"`
	UserID uint `json:"user_id" gorm:"primaryKey:autoIncrement:false"`
	Type VoteAnswerType `json:"type"`
}

