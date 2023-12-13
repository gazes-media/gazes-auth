package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetDB returns the instance of the database connection.
// If the connection is not already established, it establishes a new connection using the environment variables:
// - DATABASE_HOST: the host of the database server
// - DATABASE_PORT: the port number of the database server
// - DATABASE_USER: the username for the database connection
// - DATABASE_PASSWORD: the password for the database connection
// - DATABASE_NAME: the name of the database
// It returns a pointer to the gorm.DB object representing the database connection.
func GetDB() *gorm.DB {
	if DB == nil {
		ValidateEnvVars([]string{"DATABASE_HOST", "DATABASE_PORT", "DATABASE_USER", "DATABASE_PASSWORD", "DATABASE_NAME"})

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PORT"))

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic(err)
		}

		DB = db
	}

	return DB
}
