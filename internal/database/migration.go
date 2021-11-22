package database

import (
	"rsathishtechit/go-rest-api/internal/comment"

	"gorm.io/gorm"
)

//MigrateDB -  migrates our databse and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result != nil {
		return result
	}
	return nil
}