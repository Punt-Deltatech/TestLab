package main

import (
	"log"

	"learnhub/internal/config"
)

func main() {
	db, err := config.Connect("learnhub.db")
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	log.Println("OK: migrated learnhub.db — เปิดด้วย SQLite Viewer เพื่อตรวจ schema")
}
