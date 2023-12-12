package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
// Init initializes the database, creating a connection and a schema
func Init() {
	// we need to heck if every environment variable is set 
	// if not, we need to panic
	if os.Getenv("DATABASE_HOST") == "" {
		panic("DATABASE_HOST environment variable is not set")
	}
	if os.Getenv("DATABASE_USER") == "" {
		panic("DATABASE_USER environment variable is not set")
	}
	if os.Getenv("DATABASE_PASSWORD") == "" {
		panic("DATABASE_PASSWORD environment variable is not set")
	}
	if os.Getenv("DATABASE_NAME") == "" {
		panic("DATABASE_NAME environment variable is not set")
	}
	if os.Getenv("DATABASE_PORT") == "" {
		panic("DATABASE_PORT environment variable is not set")
	}
	// we need to create a connection string
	// the connection string is the way we connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris",os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	db.AutoMigrate(&User{})
}
