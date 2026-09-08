package main

import (
	"log"

	"libraryexam/internal/config"
)

func main() {
	db, err := config.Connect("library.db")
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	log.Println("OK: migrated library.db successfully — เปิดไฟล์ library.db ด้วย SQLite Viewer เพื่อตรวจ schema")
}
