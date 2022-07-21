package models

import (
	"github.com/jinzhu/gorm"
)

type Topic struct {
	gorm.Model
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	Questions []*Question `gorm:"many2many:topic_question"`
}
