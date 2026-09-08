package models

import "time"

type Lesson struct {
	ID      uint   `gorm:"column:lesson_id;primaryKey;autoIncrement" json:"lesson_id"`
	Title   string `gorm:"size:255;not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`

	// composite unique index: ลำดับบท (order_no) ห้ามซ้ำภายในคอร์สเดียวกัน
	CourseID uint `gorm:"column:course_id;not null;uniqueIndex:uq_course_order" json:"course_id"`
	OrderNo  int  `gorm:"not null;uniqueIndex:uq_course_order" json:"order_no"`

	DurationSeconds *int `gorm:"column:duration_seconds" json:"duration_seconds,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Course *Course `gorm:"foreignKey:CourseID;references:ID" json:"-"`
}

func (Lesson) TableName() string { return "lessons" }
