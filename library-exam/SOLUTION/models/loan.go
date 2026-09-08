package models

import "time"

type Loan struct {
	ID uint `gorm:"column:loan_id;primaryKey;autoIncrement" json:"loan_id"`

	BookID   uint `gorm:"column:book_id;not null" json:"book_id"`
	MemberID uint `gorm:"column:member_id;not null" json:"member_id"`

	LoanedAt   time.Time  `gorm:"not null" json:"loaned_at"`
	DueDate    time.Time  `gorm:"not null" json:"due_date"`
	ReturnedAt *time.Time `gorm:"column:returned_at" json:"returned_at,omitempty"` // nullable

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Book   *Book   `gorm:"foreignKey:BookID;references:ID" json:"book,omitempty"`
	Member *Member `gorm:"foreignKey:MemberID;references:ID" json:"member,omitempty"`
}

func (Loan) TableName() string { return "loans" }
