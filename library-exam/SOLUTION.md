# เฉลย + คำอธิบาย

โค้ดเฉลยเต็มอยู่ใน `SOLUTION/models/` — ก๊อปทับ `internal/models/` แล้วรัน
`go test ./tests/... -v` ได้ `ok` ทุกข้อ และ `go run ./cmd/server` migrate ผ่าน

| # | ไฟล์ / ฟิลด์ | อาการเดิม | แก้เป็น | หลักการ |
|---|---|---|---|---|
| 1 | `category.go` `ID` | ไม่มี tag เลย → GORM เดาเป็น PK ชื่อ `id` | `gorm:"column:category_id;primaryKey;autoIncrement"` | ตั้ง **Primary Key** + ชื่อคอลัมน์ตาม diagram |
| 2 | `category.go` `Name` | ไม่มี `not null`, ไม่มี unique | `gorm:"size:100;not null;uniqueIndex"` | ชื่อหมวดหมู่ห้ามซ้ำ/ห้ามว่าง → **UNIQUE + NOT NULL** |
| 3 | `member.go` `Email` | `not null` แต่ยังซ้ำได้ | เพิ่ม `uniqueIndex` | อีเมลระบุตัวตน → **UNIQUE** |
| 4 | `member.go` `Phone` | `string` + `not null` | `*string` (ตัด `not null`) | ค่าที่ "ไม่บังคับ" ต้องเป็น **pointer** เพื่อแทน `NULL` |
| 5 | `member.go` `LoanID` | Member ถือคอลัมน์ `loan_id` | ลบฟิลด์ทิ้ง | 1:N Member→Loan → ฝั่ง **"many" (Loan) เป็นผู้ถือ FK** ไม่ใช่ฝั่ง "one" |
| 6 | `book.go` `ISBN` | `not null` แต่ซ้ำได้ | เพิ่ม `uniqueIndex` | รหัสสากลประจำเล่ม → **UNIQUE** |
| 7 | `book.go` `PublishedYear` | `int` | `*int` | ปีพิมพ์ไม่บังคับ → **pointer = nullable** |
| 8 | `book.go` `CategoryID` / `Category` | `gorm:"-"` (ถูก ignore) | เอา `-` ออก, `CategoryID` ใส่ `not null`, `Category` ใส่ `foreignKey:CategoryID;references:ID` | หนังสือทุกเล่มต้องมีหมวด → **FK not null + ประกาศ belongs-to** |
| 9 | `book_author.go` | มี `ID` เดี่ยวเป็น PK | ลบ `ID`, ใส่ `primaryKey` ทั้ง `BookID` และ `AuthorID` | junction M:N + กันซ้ำ → **Composite Primary Key** |
| 10 | `loan.go` `MemberID` | `gorm:"-"` | เอา `-` ออก, ใส่ `column:member_id;not null` | Loan ต้องถือ FK ไปสมาชิก |
| 11 | `loan.go` `ReturnedAt` | `time.Time` + `not null` | `*time.Time` (ตัด `not null`) | ยังไม่คืน = `NULL` → **pointer** |
| 12 | `membership_card.go` `MemberID` | `not null` แต่ซ้ำได้ | เพิ่ม `uniqueIndex` | 1:1 บังคับด้วย **UNIQUE บนคอลัมน์ FK** |

## สรุปหลักการที่ออกสอบบ่อย

1. **ใครถือ Foreign Key** — ในความสัมพันธ์ 1:N ฝั่ง "N" (ลูก) เป็นผู้ถือ FK เสมอ
   เช่น Loan ถือ `member_id`/`book_id`, Book ถือ `category_id`
2. **1:1** = เอา FK ไปไว้ฝั่งใดฝั่งหนึ่ง แล้วทำให้คอลัมน์นั้น **UNIQUE**
3. **M:N** = ตารางเชื่อม (junction) + **Composite PK** จากคู่ FK ทั้งสอง
4. **Nullable** ใน Go = ใช้ pointer (`*string`, `*int`, `*time.Time`) และห้ามมี `not null`
5. **UNIQUE** ใช้ tag `uniqueIndex` (หรือ `unique`) — คนละเรื่องกับ `primaryKey`
6. อย่าเผลอใส่ `gorm:"-"` ทิ้งไว้ (คอลัมน์นั้นจะหายไปทั้งคอลัมน์)

## เช็คใน SQLite Viewer

- ตาราง `book_authors` ต้องมี **PK 2 คอลัมน์** (ดูที่คอลัมน์ทำเครื่องหมาย 🔑 สองอัน)
- ตาราง `members` ต้อง **ไม่มี** `loan_id`
- ตาราง `loans` มี `book_id`, `member_id` และ `returned_at` เป็น nullable
- ตาราง `membership_cards` — `member_id` มี UNIQUE
- ทุก FK ควรโชว์ในแท็บ foreign keys ของ Viewer
