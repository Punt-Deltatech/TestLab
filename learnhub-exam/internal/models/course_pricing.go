package models

import "time"

type CoursePricing struct {
	ID       uint   `gorm:"column:pricing_id;primaryKey;autoIncrement" json:"pricing_id"`
	Currency string `gorm:"size:3;not null;default:USD" json:"currency"`

	CourseID uint    `gorm:"column:course_id;not null" json:"course_id"`
	Course   *Course `gorm:"foreignKey:CourseID" json:"-"`

	DiscountPercent int      `gorm:"not null;default:0" json:"discount_percent"`
	PromoEndsAt     time.Time `gorm:"not null" json:"promo_ends_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (CoursePricing) TableName() string {
	return "course_pricings"
}
