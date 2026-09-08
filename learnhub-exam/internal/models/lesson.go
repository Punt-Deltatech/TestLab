package models

import "time"

type Lesson struct {
	ID      uint   `gorm:"column:lesson_id;primaryKey;autoIncrement" json:"lesson_id"`
	Title   string `gorm:"size:255;not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`

	CourseID uint `gorm:"column:course_id;not null" json:"course_id"`

	OrderNo         int  `gorm:"not null" json:"order_no"`
	DurationSeconds *int `gorm:"column:duration_seconds" json:"duration_seconds,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Lesson) TableName() string {
	return "lessons"
}
