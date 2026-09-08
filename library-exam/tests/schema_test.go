package tests

import (
	"reflect"
	"strings"
	"testing"

	"libraryexam/internal/models"
)

// 1. Category: table name + PK category_id + name unique & not null
func TestCategorySchema(t *testing.T) {
	sch := parse(t, &models.Category{})

	if sch.Table != "categories" {
		t.Fatalf("expected table 'categories', got %q", sch.Table)
	}

	pk := make([]string, 0)
	for _, f := range sch.PrimaryFields {
		pk = append(pk, f.DBName)
	}
	if len(pk) != 1 || pk[0] != "category_id" {
		t.Fatalf("expected single primary key 'category_id', got %v (ใส่ gorm:\"column:category_id;primaryKey\" ให้ ID)", pk)
	}

	name := sch.FieldsByDBName["name"]
	if !isNotNull(name) {
		t.Fatalf("Category.Name ต้องมี not null")
	}
	if !hasUnique(name) {
		t.Fatalf("Category.Name ต้องมี uniqueIndex (ห้ามมีชื่อหมวดหมู่ซ้ำ)")
	}
}

// 2. Author: sanity check (ไฟล์นี้ถูกมาแล้ว ควรผ่านตั้งแต่แรก)
func TestAuthorSchema(t *testing.T) {
	sch := parse(t, &models.Author{})
	if sch.Table != "authors" {
		t.Fatalf("expected table 'authors', got %q", sch.Table)
	}
	if f := sch.FieldsByDBName["biography"]; f == nil || f.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Author.Biography ควรเป็น *string (nullable)")
	}
}

// 3. Member: email unique, phone nullable pointer, ไม่ถือ loan_id
func TestMemberSchema(t *testing.T) {
	sch := parse(t, &models.Member{})

	if sch.Table != "members" {
		t.Fatalf("expected table 'members', got %q", sch.Table)
	}

	email := sch.FieldsByDBName["email"]
	if !hasUnique(email) {
		t.Fatalf("Member.Email ต้องมี uniqueIndex")
	}

	phone := sch.FieldsByDBName["phone"]
	if phone == nil {
		t.Fatalf("ไม่พบคอลัมน์ phone")
	}
	if isNotNull(phone) {
		t.Fatalf("Member.Phone ต้องเว้นว่างได้ (เอา not null ออก)")
	}
	if phone.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Member.Phone ต้องเป็น *string เพื่อแทนค่า NULL")
	}

	if f, ok := sch.FieldsByDBName["loan_id"]; ok && f.DBName != "" {
		t.Fatalf("Member ไม่ควรถือคอลัมน์ loan_id — ฝั่งที่ถือ FK คือ Loan (ลบ LoanID ออกจาก Member)")
	}
}

// 4. Book: isbn unique, published_year nullable pointer, category_id เป็น FK จริง (not null)
func TestBookSchema(t *testing.T) {
	sch := parse(t, &models.Book{})

	if sch.Table != "books" {
		t.Fatalf("expected table 'books', got %q", sch.Table)
	}

	isbn := sch.FieldsByDBName["isbn"]
	if !hasUnique(isbn) {
		t.Fatalf("Book.ISBN ต้องมี uniqueIndex")
	}

	py := sch.FieldsByDBName["published_year"]
	if py == nil {
		t.Fatalf("ไม่พบคอลัมน์ published_year")
	}
	if py.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Book.PublishedYear ต้องเป็น *int (nullable)")
	}

	cat := sch.FieldsByDBName["category_id"]
	if cat == nil || cat.DBName == "" {
		t.Fatalf("Book ต้องมีคอลัมน์ category_id (เอา gorm:\"-\" ออก)")
	}
	if !isNotNull(cat) {
		t.Fatalf("Book.CategoryID ต้องเป็น not null (หนังสือทุกเล่มต้องมีหมวดหมู่)")
	}
	if _, ok := sch.Relationships.Relations["Category"]; !ok {
		t.Fatalf("ต้องประกาศความสัมพันธ์ Book belongs to Category")
	}
}

// 5. BookAuthor: composite primary key (book_id, author_id)
func TestBookAuthorCompositePK(t *testing.T) {
	sch := parse(t, &models.BookAuthor{})

	if sch.Table != "book_authors" {
		t.Fatalf("expected table 'book_authors', got %q", sch.Table)
	}

	pk := make([]string, 0)
	for _, f := range sch.PrimaryFields {
		pk = append(pk, strings.ToLower(f.DBName))
	}
	if len(pk) != 2 {
		t.Fatalf("BookAuthor ต้องมี Composite PK 2 คอลัมน์ ได้ %d: %v (ลบ ID ออก ใส่ primaryKey ให้ BookID+AuthorID)", len(pk), pk)
	}
	got := strings.Join(pk, ",")
	if !strings.Contains(got, "book_id") || !strings.Contains(got, "author_id") {
		t.Fatalf("Composite PK ต้องเป็น (book_id, author_id), ได้: %v", pk)
	}
}

// 6. Loan: ถือ book_id + member_id, returned_at nullable pointer
func TestLoanSchema(t *testing.T) {
	sch := parse(t, &models.Loan{})

	if sch.Table != "loans" {
		t.Fatalf("expected table 'loans', got %q", sch.Table)
	}

	for _, col := range []string{"book_id", "member_id"} {
		f, ok := sch.FieldsByDBName[col]
		if !ok || f.DBName == "" {
			t.Fatalf("Loan ต้องมีคอลัมน์ %s (เอา gorm:\"-\" ออกถ้ามี)", col)
		}
		if !isNotNull(f) {
			t.Fatalf("Loan.%s ควรเป็น not null", col)
		}
	}

	ret := sch.FieldsByDBName["returned_at"]
	if ret == nil {
		t.Fatalf("ไม่พบคอลัมน์ returned_at")
	}
	if isNotNull(ret) {
		t.Fatalf("Loan.ReturnedAt ต้องเว้นว่างได้ (ยังไม่คืน = NULL)")
	}
	if ret.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Loan.ReturnedAt ต้องเป็น *time.Time")
	}
}

// 7. MembershipCard: member_id uniqueIndex -> 1:1
func TestMembershipCardOneToOne(t *testing.T) {
	sch := parse(t, &models.MembershipCard{})

	if sch.Table != "membership_cards" {
		t.Fatalf("expected table 'membership_cards', got %q", sch.Table)
	}

	mid := sch.FieldsByDBName["member_id"]
	if mid == nil || mid.DBName == "" {
		t.Fatalf("ไม่พบคอลัมน์ member_id")
	}
	if !hasUnique(mid) {
		t.Fatalf("MembershipCard.MemberID ต้องมี uniqueIndex เพื่อบังคับ 1:1 กับ Member")
	}
}
