package models

import "time"

// Enrollment = junction M:N (User[student] x Course), composite PK กันลงทะเบียนซ้ำ
type Enrollment struct {
	StudentID uint `gorm:"column:student_id;primaryKey" json:"student_id"`
	CourseID  uint `gorm:"column:course_id;primaryKey" json:"course_id"`

	EnrolledAt      time.Time  `gorm:"not null" json:"enrolled_at"`
	ProgressPercent int        `gorm:"not null;default:0" json:"progress_percent"`
	CompletedAt     *time.Time `gorm:"column:completed_at" json:"completed_at,omitempty"` // NULL = ยังเรียนไม่จบ

	Student *User   `gorm:"foreignKey:StudentID;references:ID" json:"-"`
	Course  *Course `gorm:"foreignKey:CourseID;references:ID" json:"course,omitempty"`
}

func (Enrollment) TableName() string { return "enrollments" }
