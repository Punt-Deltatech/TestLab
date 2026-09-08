package models

// BookAuthor — ตารางเชื่อม M:N ระหว่าง Book กับ Author
// หนังสือ 1 เล่มมีผู้แต่งได้หลายคน และผู้แต่ง 1 คนมีหนังสือได้หลายเล่ม
type BookAuthor struct {
	BookID   uint `gorm:"column:book_id;primaryKey" json:"book_id"`
	AuthorID uint `gorm:"column:author_id;primaryKey" json:"author_id"`

	// ลำดับ/บทบาทของผู้แต่ง เช่น "main", "co-author", "editor" (ไม่บังคับ)
	AuthorRole *string `gorm:"column:author_role;size:50" json:"author_role,omitempty"`

	Book   *Book   `gorm:"foreignKey:BookID;references:ID" json:"-"`
	Author *Author `gorm:"foreignKey:AuthorID;references:ID" json:"author,omitempty"`
}

func (BookAuthor) TableName() string {
	return "book_authors"
}

// -----------------------------------------------------------------------------
// TODO (BUG #9): ระบบห้ามบันทึกผู้แต่งคนเดิมซ้ำในหนังสือเล่มเดียวกัน
//                จึงต้องใช้ (book_id, author_id) ร่วมกันเป็น Composite Primary Key
//                -> ลบฟิลด์ ID ทิ้ง แล้วใส่ primaryKey ให้ทั้ง BookID และ AuthorID
// -----------------------------------------------------------------------------
