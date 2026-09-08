package models

import "time"

type Enrollment struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	StudentID uint `gorm:"column:student_id;not null" json:"student_id"`
	CourseID  uint `gorm:"column:course_id;not null" json:"course_id"`

	EnrolledAt      time.Time `gorm:"not null" json:"enrolled_at"`
	ProgressPercent int       `gorm:"not null;default:0" json:"progress_percent"`
	CompletedAt     time.Time `gorm:"not null" json:"completed_at"`

	Student *User   `gorm:"foreignKey:StudentID" json:"-"`
	Course  *Course `gorm:"foreignKey:CourseID" json:"course,omitempty"`
}

func (Enrollment) TableName() string {
	return "enrollments"
}
