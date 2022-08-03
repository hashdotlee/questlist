package models 

import (
	"github.com/jinzhu/gorm"
)

type UserRole int

	const (
		UserRoleAdmin UserRole = iota
		UserRoleCommon
	)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `gorm:"unique" json:"email"`
	Role UserRole ` json:"role"`
	Phone string `json:"phone"`
	VoteAnswer []VoteAnswer `json:"voteAnswer"`
	VoteQuestion []VoteQuestion `json:"voteQuestion"`
	Birthday string `json:"birthday"`
	Verified bool `json:"verified"`
	Address string `json:"address"`
}
