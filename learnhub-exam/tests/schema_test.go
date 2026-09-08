package tests

import (
	"reflect"
	"strings"
	"testing"

	"learnhub/internal/models"
)

// ---------------------------------------------------------------------------
// 1. Category — self-referential hierarchy
// ---------------------------------------------------------------------------
func TestCategory(t *testing.T) {
	sch := parse(t, &models.Category{})

	if sch.Table != "categories" {
		t.Fatalf("table ต้องเป็น 'categories' ได้ %q", sch.Table)
	}
	if pk := pkNames(sch); len(pk) != 1 || pk[0] != "category_id" {
		t.Fatalf("PK ต้องเป็นคอลัมน์เดียวชื่อ 'category_id' ได้ %v", pk)
	}
	if !hasUnique(field(t, sch, "slug")) {
		t.Fatalf("Category.Slug ต้อง unique")
	}
	parent := field(t, sch, "parent_id")
	if isNotNull(parent) {
		t.Fatalf("Category.ParentID ต้องเป็น NULL ได้ (หมวดระดับบนสุดไม่มี parent)")
	}
	if parent.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Category.ParentID ต้องเป็น *uint")
	}
	if _, ok := sch.Relationships.Relations["Parent"]; !ok {
		t.Fatalf("ต้องมีความสัมพันธ์ self-reference 'Parent' (Category -> Category)")
	}
}

// ---------------------------------------------------------------------------
// 2. User — Single Table Inheritance (instructor / student)
// ---------------------------------------------------------------------------
func TestUserSTI(t *testing.T) {
	sch := parse(t, &models.User{})

	if sch.Table != "users" {
		t.Fatalf("table ต้องเป็น 'users' ได้ %q", sch.Table)
	}
	if !hasUnique(field(t, sch, "email")) {
		t.Fatalf("User.Email ต้อง unique")
	}

	// ต้องมีคอลัมน์ discriminator อย่างใดอย่างหนึ่ง
	disc := false
	for _, c := range []string{"user_type", "role", "type"} {
		if f, ok := sch.FieldsByDBName[c]; ok && f.DBName != "" {
			disc = true
		}
	}
	if !disc {
		t.Fatalf("ต้องมีคอลัมน์แยกประเภทผู้ใช้ เช่น 'user_type'")
	}

	// subtype fields: ต้อง nullable + pointer ทั้งหมด
	subtypes := []string{"headline", "years_experience", "study_goal"}
	for _, name := range subtypes {
		f := field(t, sch, name)
		if isNotNull(f) {
			t.Fatalf("User.%s เป็นข้อมูลเฉพาะบทบาท ต้อง NULL ได้ (เอา not null ออก)", f.Name)
		}
		if f.FieldType.Kind() != reflect.Ptr {
			t.Fatalf("User.%s ต้องเป็น pointer (เช่น *string / *int) ใน STI", f.Name)
		}
	}

	// mentor self-reference (optional)
	mentor := field(t, sch, "mentor_id")
	if isNotNull(mentor) || mentor.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("User.MentorID ต้องเป็น *uint และ NULL ได้ (ผู้ใช้ไม่จำเป็นต้องมี mentor)")
	}
	if _, ok := sch.Relationships.Relations["Mentor"]; !ok {
		t.Fatalf("ต้องมีความสัมพันธ์ self-reference 'Mentor' (User -> User)")
	}
}

// ---------------------------------------------------------------------------
// 3. Course — belongs to Instructor(User) + Category, publish date optional
// ---------------------------------------------------------------------------
func TestCourse(t *testing.T) {
	sch := parse(t, &models.Course{})

	if sch.Table != "courses" {
		t.Fatalf("table ต้องเป็น 'courses' ได้ %q", sch.Table)
	}

	for _, col := range []string{"instructor_id", "category_id"} {
		f := field(t, sch, col)
		if !isNotNull(f) {
			t.Fatalf("Course.%s ต้อง not null (คอร์สต้องมีผู้สอนและหมวดหมู่เสมอ)", f.Name)
		}
	}
	for _, rel := range []string{"Instructor", "Category"} {
		if _, ok := sch.Relationships.Relations[rel]; !ok {
			t.Fatalf("Course ต้องประกาศความสัมพันธ์ belongs-to %q (อย่าใช้ gorm:\"-\")", rel)
		}
	}

	pub := field(t, sch, "published_at")
	if isNotNull(pub) {
		t.Fatalf("Course.PublishedAt ต้อง NULL ได้ (คอร์สที่ยังเป็น draft ยังไม่มีวันเผยแพร่)")
	}
	if pub.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Course.PublishedAt ต้องเป็น *time.Time")
	}
}

// ---------------------------------------------------------------------------
// 4. CoursePricing — 1:1 optional กับ Course
// ---------------------------------------------------------------------------
func TestCoursePricingOneToOne(t *testing.T) {
	sch := parse(t, &models.CoursePricing{})

	if sch.Table != "course_pricings" {
		t.Fatalf("table ต้องเป็น 'course_pricings' ได้ %q", sch.Table)
	}
	cid := field(t, sch, "course_id")
	if !isNotNull(cid) {
		t.Fatalf("CoursePricing.CourseID ต้อง not null")
	}
	if !hasUnique(cid) {
		t.Fatalf("CoursePricing.CourseID ต้อง unique เพื่อบังคับ 1:1 (คอร์สมี pricing ได้ไม่เกิน 1)")
	}
}

// ---------------------------------------------------------------------------
// 5. Lesson — 1:N กับ Course + ลำดับบท (order_no) ห้ามซ้ำในคอร์สเดียว
// ---------------------------------------------------------------------------
func TestLessonCompositeUnique(t *testing.T) {
	sch := parse(t, &models.Lesson{})

	if sch.Table != "lessons" {
		t.Fatalf("table ต้องเป็น 'lessons' ได้ %q", sch.Table)
	}
	_ = field(t, sch, "course_id")
	_ = field(t, sch, "order_no")

	// ต้องมี composite UNIQUE index ที่คลุมทั้ง course_id และ order_no
	found := false
	for _, idx := range sch.ParseIndexes() {
		if strings.ToUpper(idx.Class) != "UNIQUE" {
			continue
		}
		cols := map[string]bool{}
		for _, opt := range idx.Fields {
			if opt.Field != nil {
				cols[opt.Field.DBName] = true
			}
		}
		if cols["course_id"] && cols["order_no"] {
			found = true
		}
	}
	if !found {
		t.Fatalf("ต้องมี composite uniqueIndex คลุม (course_id, order_no) เพื่อกันลำดับบทซ้ำในคอร์สเดียว\n" +
			"เช่น: `gorm:\"...;uniqueIndex:uq_course_order\"` ที่ทั้งสองฟิลด์")
	}

	if d := field(t, sch, "duration_seconds"); d.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Lesson.DurationSeconds ต้องเป็น *int (บางบทยังไม่ระบุความยาว)")
	}
}

// ---------------------------------------------------------------------------
// 6. Enrollment — junction M:N (student x course) + composite PK
// ---------------------------------------------------------------------------
func TestEnrollmentCompositePK(t *testing.T) {
	sch := parse(t, &models.Enrollment{})

	if sch.Table != "enrollments" {
		t.Fatalf("table ต้องเป็น 'enrollments' ได้ %q", sch.Table)
	}
	pk := pkNames(sch)
	if len(pk) != 2 || !contains(pk, "student_id") || !contains(pk, "course_id") {
		t.Fatalf("Enrollment ต้องมี Composite PK = (student_id, course_id) ได้ %v", pk)
	}
	c := field(t, sch, "completed_at")
	if isNotNull(c) || c.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Enrollment.CompletedAt ต้องเป็น *time.Time และ NULL ได้ (ยังเรียนไม่จบ)")
	}
}

// ---------------------------------------------------------------------------
// 7. Review — junction M:N + composite PK, comment optional
// ---------------------------------------------------------------------------
func TestReviewCompositePK(t *testing.T) {
	sch := parse(t, &models.Review{})

	if sch.Table != "reviews" {
		t.Fatalf("table ต้องเป็น 'reviews' ได้ %q", sch.Table)
	}
	pk := pkNames(sch)
	if len(pk) != 2 || !contains(pk, "student_id") || !contains(pk, "course_id") {
		t.Fatalf("Review ต้องมี Composite PK = (student_id, course_id) — รีวิวได้คอร์สละ 1 ครั้ง ได้ %v", pk)
	}
	cm := field(t, sch, "comment")
	if isNotNull(cm) {
		t.Fatalf("Review.Comment ต้อง NULL ได้ (ให้ดาวเฉย ๆ ไม่พิมพ์ข้อความก็ได้)")
	}
	if cm.FieldType.Kind() != reflect.Ptr {
		t.Fatalf("Review.Comment ต้องเป็น *string")
	}
}
