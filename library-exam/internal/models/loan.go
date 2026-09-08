package models

import "time"

// Loan — รายการยืมหนังสือ 1 ครั้ง (ยืมหนังสือ 1 เล่ม โดยสมาชิก 1 คน)
type Loan struct {
	ID uint `gorm:"column:loan_id;primaryKey;autoIncrement" json:"loan_id"`

	BookID   uint `gorm:"column:book_id;not null" json:"book_id"`
	MemberID uint `gorm:"column:member_id;not null" json:"member_id"`

	LoanedAt time.Time `gorm:"not null" json:"loaned_at"`
	DueDate  time.Time `gorm:"not null" json:"due_date"`

	// เวลาที่คืนหนังสือ — ถ้ายังไม่คืนต้องเป็น NULL
	ReturnedAt *time.Time `gorm:"" json:"returned_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Book   *Book   `gorm:"foreignKey:BookID;references:ID" json:"book,omitempty"`
	Member *Member `gorm:"foreignKey:MemberID;references:ID" json:"member,omitempty"`
}

func (Loan) TableName() string {
	return "loans"
}

// -----------------------------------------------------------------------------
// TODO (BUG #10): Loan เป็นฝั่งที่ต้องถือ FK member_id (1:N Member->Loan)
//                 ตอนนี้ MemberID ถูก ignore ด้วย gorm:"-" -> เอาออก แล้วใส่ not null
// TODO (BUG #11): ReturnedAt ต้องเป็น NULL ได้ (ยืมอยู่ยังไม่คืน)
//                 -> เปลี่ยนเป็น *time.Time และเอา not null ออก
// -----------------------------------------------------------------------------
