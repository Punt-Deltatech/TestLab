package models

import "time"

// Review = junction M:N (User[student] x Course), composite PK -> รีวิวได้คอร์สละ 1 ครั้ง
type Review struct {
	StudentID uint `gorm:"column:student_id;primaryKey" json:"student_id"`
	CourseID  uint `gorm:"column:course_id;primaryKey" json:"course_id"`

	Rating  int     `gorm:"not null" json:"rating"`
	Comment *string `gorm:"size:1000" json:"comment,omitempty"` // NULL = ให้ดาวเฉย ๆ

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Student *User   `gorm:"foreignKey:StudentID;references:ID" json:"-"`
	Course  *Course `gorm:"foreignKey:CourseID;references:ID" json:"-"`
}

func (Review) TableName() string { return "reviews" }
