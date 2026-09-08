package models

import "time"

type Course struct {
	ID       uint   `gorm:"column:course_id;primaryKey;autoIncrement" json:"course_id"`
	Title    string `gorm:"size:255;not null" json:"title"`
	Subtitle string `gorm:"size:255" json:"subtitle"`

	InstructorID uint  `gorm:"column:instructor_id" json:"instructor_id"`
	Instructor   *User `gorm:"-" json:"instructor,omitempty"`

	CategoryID uint      `gorm:"-" json:"category_id"`
	Category   *Category `gorm:"-" json:"category,omitempty"`

	PriceCents  int       `gorm:"not null;default:0" json:"price_cents"`
	PublishedAt time.Time `gorm:"not null" json:"published_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Lessons     []Lesson     `gorm:"foreignKey:CourseID" json:"lessons,omitempty"`
	Enrollments []Enrollment `json:"enrollments,omitempty"`
	Pricing     *CoursePricing `json:"pricing,omitempty"`
}

func (Course) TableName() string {
	return "courses"
}
