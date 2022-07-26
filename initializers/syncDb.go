package initializers

import (
	"dblab/questlist/models"
)
func SyncDb(){
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Topic{})
	DB.AutoMigrate(&models.Topic{}, &models.Question{})
	DB.AutoMigrate(&models.Question{})
	DB.AutoMigrate(&models.Answer{})
	DB.AutoMigrate(&models.VoteAnswer{})
	DB.AutoMigrate(&models.VoteQuestion{})
}
