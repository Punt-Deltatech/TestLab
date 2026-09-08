package models

import "time"

type Review struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	StudentID uint `gorm:"column:student_id;not null" json:"student_id"`
	CourseID  uint `gorm:"column:course_id;not null" json:"course_id"`

	Rating  int    `gorm:"not null" json:"rating"`
	Comment string `gorm:"size:1000;not null" json:"comment"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Student *User   `gorm:"foreignKey:StudentID" json:"-"`
	Course  *Course `gorm:"foreignKey:CourseID" json:"-"`
}

func (Review) TableName() string {
	return "reviews"
}
