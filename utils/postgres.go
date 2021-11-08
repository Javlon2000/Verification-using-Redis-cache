package utils

import (
	"os"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"github.com/joho/godotenv"
)

func DB() (*gorm.DB, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Cannot load the .env: %v", err)
	}

	DB_HOST := os.Getenv("APP_DB_HOST")
	DB_USER := os.Getenv("APP_DB_USER")
	DB_PASSWORD := os.Getenv("APP_DB_PASSWORD")
	DB_PORT := os.Getenv("APP_DB_PORT")
	DB_NAME := os.Getenv("APP_DB_NAME")
	DB_SSLMODE := os.Getenv("APP_DB_SSLMODE")
	DB_TIMEZONE := os.Getenv("APP_DB_TIMEZONE")

	DBConnection := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_PORT,
		DB_NAME,
		DB_SSLMODE,
		DB_TIMEZONE,
	)

	return gorm.Open(postgres.Open(DBConnection), &gorm.Config{})
}

