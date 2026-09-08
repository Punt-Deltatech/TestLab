package models

import "time"

type User struct {
	ID       uint   `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	FullName string `gorm:"size:150;not null" json:"full_name"`
	Email    string `gorm:"size:255;not null" json:"email"`
	UserType string `gorm:"size:20;not null" json:"user_type"`

	Headline        string `gorm:"size:200;not null" json:"headline"`
	YearsExperience int    `gorm:"not null" json:"years_experience"`

	StudyGoal string `gorm:"size:255;not null" json:"study_goal"`

	MentorID uint  `gorm:"column:mentor_id;not null" json:"mentor_id"`
	Mentor   *User `gorm:"foreignKey:MentorID" json:"mentor,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	CoursesTaught []Course     `gorm:"foreignKey:InstructorID" json:"courses_taught,omitempty"`
	Enrollments   []Enrollment `json:"enrollments,omitempty"`
}

func (User) TableName() string {
	return "users"
}
