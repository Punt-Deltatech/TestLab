package models

import "time"

type MembershipCard struct {
	ID         uint   `gorm:"column:card_id;primaryKey;autoIncrement" json:"card_id"`
	CardNumber string `gorm:"size:50;not null;uniqueIndex" json:"card_number"`

	MemberID uint    `gorm:"column:member_id;not null;uniqueIndex" json:"member_id"` // unique -> 1:1
	Member   *Member `gorm:"foreignKey:MemberID;references:ID" json:"-"`

	IssuedAt  time.Time `gorm:"not null" json:"issued_at"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (MembershipCard) TableName() string { return "membership_cards" }
