package models

import "time"

type Course struct {
	ID       uint   `gorm:"column:course_id;primaryKey;autoIncrement" json:"course_id"`
	Title    string `gorm:"size:255;not null" json:"title"`
	Subtitle string `gorm:"size:255" json:"subtitle"`

	InstructorID uint  `gorm:"column:instructor_id;not null" json:"instructor_id"`
	Instructor   *User `gorm:"foreignKey:InstructorID;references:ID" json:"instructor,omitempty"`

	CategoryID uint      `gorm:"column:category_id;not null" json:"category_id"`
	Category   *Category `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`

	PriceCents  int        `gorm:"not null;default:0" json:"price_cents"`
	PublishedAt *time.Time `gorm:"column:published_at" json:"published_at,omitempty"` // NULL = ยังเป็น draft

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Lessons     []Lesson       `gorm:"foreignKey:CourseID;references:ID" json:"lessons,omitempty"`
	Enrollments []Enrollment   `gorm:"foreignKey:CourseID;references:ID" json:"enrollments,omitempty"`
	Pricing     *CoursePricing `gorm:"foreignKey:CourseID;references:ID" json:"pricing,omitempty"`
}

func (Course) TableName() string { return "courses" }
