package models

import "time"

type Category struct {
	ID   uint   `gorm:"column:category_id;primaryKey;autoIncrement" json:"category_id"`
	Name string `gorm:"size:120;not null" json:"name"`
	Slug string `gorm:"size:120;not null;uniqueIndex" json:"slug"`

	ParentID *uint     `gorm:"column:parent_id" json:"parent_id,omitempty"` // NULL = หมวดระดับบนสุด
	Parent   *Category `gorm:"foreignKey:ParentID;references:ID" json:"parent,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Children []Category `gorm:"foreignKey:ParentID;references:ID" json:"children,omitempty"`
	Courses  []Course   `gorm:"foreignKey:CategoryID;references:ID" json:"courses,omitempty"`
}

func (Category) TableName() string { return "categories" }
