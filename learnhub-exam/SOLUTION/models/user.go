package models

import "time"

// User = Single Table Inheritance: instructor + student ในตาราง users เดียว
type User struct {
	ID       uint   `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	FullName string `gorm:"size:150;not null" json:"full_name"`
	Email    string `gorm:"size:255;not null;uniqueIndex" json:"email"`
	UserType string `gorm:"size:20;not null" json:"user_type"` // "instructor" | "student"

	// ---- เฉพาะ instructor (NULL เมื่อเป็น student) ----
	Headline        *string `gorm:"size:200" json:"headline,omitempty"`
	YearsExperience *int    `gorm:"column:years_experience" json:"years_experience,omitempty"`

	// ---- เฉพาะ student (NULL เมื่อเป็น instructor) ----
	StudyGoal *string `gorm:"column:study_goal;size:255" json:"study_goal,omitempty"`

	// mentor (self-reference, optional)
	MentorID *uint `gorm:"column:mentor_id" json:"mentor_id,omitempty"`
	Mentor   *User `gorm:"foreignKey:MentorID;references:ID" json:"mentor,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	CoursesTaught []Course     `gorm:"foreignKey:InstructorID;references:ID" json:"courses_taught,omitempty"`
	Enrollments   []Enrollment `gorm:"foreignKey:StudentID;references:ID" json:"enrollments,omitempty"`
}

func (User) TableName() string { return "users" }
