package config

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"learnhub/internal/models"
)

// Connect เปิด SQLite แล้ว AutoMigrate — แก้เฉพาะ internal/models/ เท่านั้น
func Connect(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := db.AutoMigrate(
		&models.Category{},
		&models.User{},
		&models.Course{},
		&models.CoursePricing{},
		&models.Lesson{},
		&models.Enrollment{},
		&models.Review{},
	); err != nil {
		return nil, fmt.Errorf("auto-migrate: %w", err)
	}
	return db, nil
}
