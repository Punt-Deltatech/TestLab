package models

import "time"

// Author — ผู้แต่งหนังสือ
type Author struct {
	ID        uint      `gorm:"column:author_id;primaryKey;autoIncrement" json:"author_id"`
	Name      string    `gorm:"size:150;not null" json:"name"`
	Biography *string   `gorm:"size:1000" json:"biography,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Author) TableName() string {
	return "authors"
}

// (ไฟล์นี้ถูกต้องแล้ว ใช้เป็นตัวอย่างอ้างอิงการเขียน GORM tags ได้)
