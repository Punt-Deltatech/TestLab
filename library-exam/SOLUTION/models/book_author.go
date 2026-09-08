package models

type BookAuthor struct {
	BookID   uint `gorm:"column:book_id;primaryKey" json:"book_id"`
	AuthorID uint `gorm:"column:author_id;primaryKey" json:"author_id"`

	AuthorRole *string `gorm:"column:author_role;size:50" json:"author_role,omitempty"`

	Book   *Book   `gorm:"foreignKey:BookID;references:ID" json:"-"`
	Author *Author `gorm:"foreignKey:AuthorID;references:ID" json:"author,omitempty"`
}

func (BookAuthor) TableName() string { return "book_authors" }
