package models 

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("sqlite3", "development.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Topic{})
	db.AutoMigrate(&Topic{}, &Question{})
	db.AutoMigrate(&Question{})
	db.AutoMigrate(&Answer{})
	db.AutoMigrate(&Contact{})

	DB = db
}
