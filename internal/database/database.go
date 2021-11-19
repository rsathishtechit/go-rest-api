package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbName, dbPassword, )
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	
	if err != nil {
		return db, nil
	}

	if _, err := db.DB(); err!=nil {
		return db, err
	}

	return db, nil
}