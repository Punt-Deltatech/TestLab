# เฉลย + คำอธิบาย (LearnHub)

โค้ดเต็มอยู่ใน `SOLUTION/models/` — ก๊อปทับ `internal/models/` แล้ว `go test ./tests/...` ได้ `ok` ทุกข้อ

## จุดที่ต้องแก้

### `category.go`
| จุด | เดิม | แก้เป็น | เหตุผล |
|---|---|---|---|
| `ID` | `gorm:"primaryKey"` → คอลัมน์ชื่อ `id` | `gorm:"column:category_id;primaryKey;autoIncrement"` | ชื่อคอลัมน์ต้องตรง diagram |
| `Slug` | ไม่มี unique | `...;uniqueIndex` | slug ใช้ทำ URL ห้ามซ้ำ |
| `ParentID` | `uint` + `not null` | `*uint` (ตัด not null) | หมวดระดับบนสุดไม่มี parent = NULL |
| `Parent` | `foreignKey:ParentID` | `foreignKey:ParentID;references:ID` | self-reference ให้ชัดเจน |

### `user.go` (STI)
| จุด | เดิม | แก้เป็น | เหตุผล |
|---|---|---|---|
| `Email` | ไม่มี unique | `...;uniqueIndex` | อีเมลห้ามซ้ำ |
| `Headline` | `string` + `not null` | `*string` | ข้อมูลเฉพาะ instructor → NULL ได้ตอนเป็น student |
| `YearsExperience` | `int` + `not null` | `*int` | เช่นเดียวกัน |
| `StudyGoal` | `string` + `not null` | `*string` | ข้อมูลเฉพาะ student → NULL ได้ตอนเป็น instructor |
| `MentorID` | `uint` + `not null` | `*uint` | mentor ไม่บังคับ |
| `Mentor` | `foreignKey:MentorID` | `foreignKey:MentorID;references:ID` | self-reference |

### `course.go`
| จุด | เดิม | แก้เป็น | เหตุผล |
|---|---|---|---|
| `InstructorID` | ไม่มี `not null` | `column:instructor_id;not null` | คอร์สต้องมีผู้สอนเสมอ |
| `Instructor` | `gorm:"-"` | `foreignKey:InstructorID;references:ID` | ต้องเป็น FK จริง ไม่ใช่ ignore |
| `CategoryID` | `gorm:"-"` | `column:category_id;not null` | คอลัมน์หายเพราะ `-` |
| `Category` | `gorm:"-"` | `foreignKey:CategoryID;references:ID` | belongs-to |
| `PublishedAt` | `time.Time` + `not null` | `*time.Time` | draft ยังไม่มีวันเผยแพร่ = NULL |

### `course_pricing.go`
| จุด | เดิม | แก้เป็น | เหตุผล |
|---|---|---|---|
| `CourseID` | `not null` เท่านั้น | `...;not null;uniqueIndex` | บังคับ 1:1 — คอร์สมี pricing ได้ไม่เกิน 1 |
| `PromoEndsAt` | `time.Time` + `not null` *(แถม)* | `*time.Time` | โปรไม่มีวันหมดอายุก็ได้ |

### `lesson.go`
| จุด | เดิม | แก้เป็น | เหตุผล |
|---|---|---|---|
| `CourseID` + `OrderNo` | ไม่มี unique ร่วมกัน | ใส่ `uniqueIndex:uq_course_order` **ที่ทั้งสองฟิลด์** | ลำดับบทห้ามซ้ำในคอร์สเดียว (composite unique index) |
| `Course` | ไม่มี relation | เพิ่ม `Course *Course gorm:"foreignKey:CourseID;references:ID"` | navigability |

> `uniqueIndex:ชื่อเดียวกัน` บนหลายฟิลด์ = GORM สร้าง `CREATE UNIQUE INDEX ... (course_id, order_no)` ให้

### `enrollment.go`
| จุด | เดิม | แก้เป็น | เหตุผล |
|---|---|---|---|
| `ID` | มี surrogate PK เดี่ยว | **ลบทิ้ง** |  |
| `StudentID`, `CourseID` | ไม่ใช่ PK | ใส่ `primaryKey` ทั้งคู่ | Composite PK กันลงทะเบียนซ้ำ (M:N junction) |
| `CompletedAt` | `time.Time` + `not null` | `*time.Time` | ยังเรียนไม่จบ = NULL |

### `review.go`
| จุด | เดิม | แก้เป็น | เหตุผล |
|---|---|---|---|
| `ID` | surrogate PK เดี่ยว | **ลบทิ้ง** |  |
| `StudentID`, `CourseID` | ไม่ใช่ PK | ใส่ `primaryKey` ทั้งคู่ | รีวิวได้คอร์สละ 1 ครั้ง |
| `Comment` | `string` + `not null` | `*string` | ให้ดาวเฉย ๆ ไม่พิมพ์ก็ได้ |

---

## หลักการที่ทดสอบในชุดนี้ (ยากกว่าชุดห้องสมุด)

1. **STI (Single Table Inheritance)** — ฟิลด์เฉพาะบทบาททุกตัวต้องเป็น pointer + ห้าม `not null`
2. **Self-referential FK** — `User.mentor_id → users`, `Category.parent_id → categories` (ต้อง `*uint` เพราะ optional + ใส่ `references:ID`)
3. **belongs-to ซ้อนกัน 2 เส้นในตารางเดียว** — `Course` ถือทั้ง `instructor_id` และ `category_id`
4. **1:1 optional** — FK + `uniqueIndex` บนตารางฝั่งลูก (`CoursePricing.course_id`)
5. **Composite UNIQUE index** (คนละเรื่องกับ composite PK) — `lessons(course_id, order_no)` ผ่าน `uniqueIndex:ชื่อเดียวกัน`
6. **Composite PRIMARY KEY** ใน junction M:N — `enrollments`, `reviews` (ลบ surrogate `ID` ก่อน)
7. **Nullable = pointer** — `*string`, `*int`, `*time.Time` และห้ามมี `not null`
8. กับดัก `gorm:"-"` ทำให้คอลัมน์หายทั้งคอลัมน์

## ตรวจใน SQLite Viewer

```
users            : email UNIQUE, mentor_id FK→users, subtype cols nullable
categories       : slug UNIQUE, parent_id FK→categories (nullable)
courses          : instructor_id + category_id NOT NULL + FK, published_at nullable
course_pricings  : course_id UNIQUE + FK  (1:1)
lessons          : UNIQUE(course_id, order_no)
enrollments      : PRIMARY KEY (student_id, course_id)  ← 🔑 สองอัน
reviews          : PRIMARY KEY (student_id, course_id)  ← 🔑 สองอัน
```
