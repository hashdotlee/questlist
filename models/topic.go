package models

import (
	"github.com/jinzhu/gorm"
)

type Topic struct {
	gorm.Model
	Title string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Questions []*Question `gorm:"many2many:topic_question" json:"questions"`
}
