package models

import "time"

type Book struct {
	ID    uint   `gorm:"column:book_id;primaryKey;autoIncrement" json:"book_id"`
	Title string `gorm:"size:255;not null" json:"title"`
	ISBN  string `gorm:"size:20;not null;uniqueIndex" json:"isbn"`

	PublishedYear *int `gorm:"column:published_year" json:"published_year,omitempty"` // nullable
	CopiesTotal   int  `gorm:"not null;default:1" json:"copies_total"`

	CategoryID uint      `gorm:"column:category_id;not null" json:"category_id"`
	Category   *Category `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Authors []BookAuthor `gorm:"foreignKey:BookID;references:ID" json:"authors,omitempty"`
	Loans   []Loan       `gorm:"foreignKey:BookID;references:ID" json:"loans,omitempty"`
}

func (Book) TableName() string { return "books" }
