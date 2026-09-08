package models

import "time"

// MembershipCard — บัตรสมาชิก
// สมาชิก 1 คนมีบัตรได้ไม่เกิน 1 ใบ (1:1 แบบ optional) และ 1 บัตรเป็นของสมาชิกคนเดียว
type MembershipCard struct {
	ID         uint   `gorm:"column:card_id;primaryKey;autoIncrement" json:"card_id"`
	CardNumber string `gorm:"size:50;not null;uniqueIndex" json:"card_number"`

	// FK ไปยังสมาชิกเจ้าของบัตร
	MemberID uint    `gorm:"column:member_id;not null;uniqueIndex" json:"member_id"`
	Member   *Member `gorm:"foreignKey:MemberID;references:ID" json:"-"`

	IssuedAt  time.Time `gorm:"not null" json:"issued_at"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (MembershipCard) TableName() string {
	return "membership_cards"
}

// -----------------------------------------------------------------------------
// TODO (BUG #12): ตอนนี้ member_id ยังไม่ unique ทำให้สมาชิก 1 คนมีบัตรได้หลายใบ
//                 -> เพิ่ม uniqueIndex ให้ MemberID เพื่อบังคับความสัมพันธ์ 1:1
// -----------------------------------------------------------------------------
