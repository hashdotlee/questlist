package models

import (
	"github.com/jinzhu/gorm"
)

type VoteAnswerType int

const (
	UpVoteAnswer VoteAnswerType = iota
	DownVoteAnswer
)

type VoteAnswer struct {
	gorm.Model
	AnswerID uint `gorm:"primaryKey" json:"answer_id"`
	UserID uint `json:"user_id" gorm:"primaryKey"`
	Type VoteAnswerType `json:"type"`
}

