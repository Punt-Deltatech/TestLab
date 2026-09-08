package models

import "time"

// Book — หนังสือในระบบ
type Book struct {
	ID    uint   `gorm:"column:book_id;primaryKey;autoIncrement" json:"book_id"`
	Title string `gorm:"size:255;not null" json:"title"`
	ISBN  string `gorm:"uniqueIndex;size:20;not null" json:"isbn"`

	// ปีที่พิมพ์ — หนังสือเก่าบางเล่มไม่ทราบปีพิมพ์
	PublishedYear *int `gorm:"" json:"published_year"`

	// จำนวนสำเนาทั้งหมดของหนังสือเล่มนี้
	CopiesTotal int `gorm:"column:copies_total;not null;" json:"copies_total"`

	// FK ไปยังหมวดหมู่ (หนังสือทุกเล่มต้องมีหมวดหมู่)
	CategoryID uint      `gorm:"column:category_id;not null" json:"category_id"`
	Category   *Category `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Authors []BookAuthor `gorm:"foreignKey:BookID;references:ID" json:"authors,omitempty"`
	Loans   []Loan       `gorm:"foreignKey:BookID;references:ID" json:"loans,omitempty"`
}

func (Book) TableName() string {
	return "books"
}

// -----------------------------------------------------------------------------
// TODO (BUG #6): ISBN เป็นรหัสสากลประจำหนังสือ ต้องไม่ซ้ำ -> เพิ่ม uniqueIndex
// TODO (BUG #7): PublishedYear ต้องเป็น NULL ได้ -> เปลี่ยนเป็น *int และเอาออกจากการบังคับค่า
// TODO (BUG #8): CategoryID / Category ตอนนี้ถูก ignore ด้วย gorm:"-"
//                ทำให้ไม่มีคอลัมน์ category_id และไม่มีความสัมพันธ์
//                -> เอา gorm:"-" ออก, ใส่ not null ให้ CategoryID,
//                   และตั้ง foreignKey/references ให้ Category ให้ถูก
// -----------------------------------------------------------------------------
