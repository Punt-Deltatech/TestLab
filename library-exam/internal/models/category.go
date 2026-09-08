package models

import "time"

// Category — หมวดหมู่ของหนังสือ
type Category struct {
	ID        uint      `gorm:"column:category_id;primaryKey;autoIncrement" json:"category_id"`
	Name      string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Books []Book `gorm:"foreignKey:CategoryID;references:ID" json:"books,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

// -----------------------------------------------------------------------------
// TODO (BUG #1): ยังไม่ได้กำหนด Primary Key ให้ ID (ต้องเป็น category_id)
// TODO (BUG #2): Name เป็นชื่อหมวดหมู่ ห้ามซ้ำ และห้ามว่าง
//                -> ต้องมี not null และ unique/uniqueIndex
// TODO: ผูกความสัมพันธ์ Category 1:N Book ให้ครบ (foreignKey อยู่ที่ Book.CategoryID)
// -----------------------------------------------------------------------------
