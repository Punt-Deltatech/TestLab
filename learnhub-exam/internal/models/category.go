package models

import "time"

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"category_id"`
	Name string `gorm:"size:120;not null" json:"name"`
	Slug string `gorm:"size:120;not null" json:"slug"`

	ParentID uint      `gorm:"column:parent_id;not null" json:"parent_id"`
	Parent   *Category `gorm:"foreignKey:ParentID" json:"parent,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Courses []Course `json:"courses,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
