package models

import "time"

type Member struct {
	ID       uint    `gorm:"column:member_id;primaryKey;autoIncrement" json:"member_id"`
	FullName string  `gorm:"size:150;not null" json:"full_name"`
	Email    string  `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Phone    *string `gorm:"size:50" json:"phone,omitempty"` // nullable

	JoinedAt time.Time `gorm:"autoCreateTime" json:"joined_at"`

	Loans          []Loan          `gorm:"foreignKey:MemberID;references:ID" json:"loans,omitempty"`
	MembershipCard *MembershipCard `gorm:"foreignKey:MemberID;references:ID" json:"membership_card,omitempty"`
}

func (Member) TableName() string { return "members" }
