package models

import "time"

type Category struct {
	ID        uint      `gorm:"column:category_id;primaryKey;autoIncrement" json:"category_id"`
	Name      string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Books []Book `gorm:"foreignKey:CategoryID;references:ID" json:"books,omitempty"`
}

func (Category) TableName() string { return "categories" }
