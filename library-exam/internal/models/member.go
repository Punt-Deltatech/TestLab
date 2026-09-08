package models

import "time"

// Member — สมาชิกห้องสมุด
type Member struct {
	ID       uint   `gorm:"column:member_id;primaryKey;autoIncrement" json:"member_id"`
	FullName string `gorm:"size:150;not null" json:"full_name"`
	Email    string `gorm:"size:255;not null;uniqueIndex" json:"email"`

	// เบอร์โทรศัพท์ — สมาชิกบางคนไม่ได้ให้เบอร์ไว้
	Phone *string `gorm:"size:50" json:"phone"`

	JoinedAt time.Time `gorm:"autoCreateTime" json:"joined_at"`

	// ความสัมพันธ์
	Loans []Loan `gorm:"foreignKey:MemberID;references:ID" json:"loans,omitempty"`
}

func (Member) TableName() string {
	return "members"
}

// -----------------------------------------------------------------------------
// TODO (BUG #3): Email ใช้ล็อกอิน/ติดต่อ ต้องไม่ซ้ำกัน -> เพิ่ม uniqueIndex
// TODO (BUG #4): Phone ต้องเว้นว่าง (NULL) ได้ -> เปลี่ยนเป็น *string และเอา not null ออก
// TODO (BUG #5): Member ไม่ควร "เป็นเจ้าของ" คอลัมน์ loan_id
//                ในความสัมพันธ์ 1:N ระหว่าง Member กับ Loan ฝั่งที่ถือ FK คือ Loan
//                -> ลบฟิลด์ LoanID ออกจาก Member
// -----------------------------------------------------------------------------
