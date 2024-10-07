package storage

import (
	"task-management-backend/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		return nil, err
	}

	return db, nil
}