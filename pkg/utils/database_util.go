package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	if DB == nil {
		vars, err := Getenvars([]string{"DATABASE_HOST", "DATABASE_PORT", "DATABASE_USER", "DATABASE_PASSWORD", "DATABASE_NAME"})
		if err != nil {
			log.Fatal(err.Error())
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris", vars["DATABASE_HOST"], vars["DATABASE_USER"], vars["DATABASE_PASSWORD"], vars["DATABASE_NAME"], vars["DATABASE_PORT"])

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic(err)
		}

		DB = db
	}

	return DB
}
