package utils

import (
	"fmt"
	"os"

	"github.com/mrspec7er/balky/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	credentials := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"))

	connection, err := gorm.Open(postgres.Open(credentials), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	Migration(connection)

	DB = connection
}

func Migration(db *gorm.DB)  {
	db.AutoMigrate(
		&model.Application{},
		&model.User{},
		&model.ReportMaster{},
		&model.Attribute{},
		&model.Content{},
	)
}
