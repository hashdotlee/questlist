package initializers 

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
)

var DB *gorm.DB

func ConnectDB() {

	// config dns by env
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// config dns
	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname

	// connect to db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}


	DB = db
}
