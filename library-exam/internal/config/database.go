package config

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"libraryexam/internal/models"
)

// Connect เปิดไฟล์ SQLite (สร้างใหม่ถ้ายังไม่มี) แล้วรัน AutoMigrate
// นักศึกษาไม่ต้องแก้ไฟล์นี้ — แก้เฉพาะ internal/models/
func Connect(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := db.AutoMigrate(
		&models.Category{},
		&models.Author{},
		&models.Member{},
		&models.Book{},
		&models.BookAuthor{},
		&models.Loan{},
		&models.MembershipCard{},
	); err != nil {
		return nil, fmt.Errorf("auto-migrate: %w", err)
	}

	return db, nil
}
