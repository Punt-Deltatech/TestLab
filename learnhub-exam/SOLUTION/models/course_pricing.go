package models

import "time"

type CoursePricing struct {
	ID       uint   `gorm:"column:pricing_id;primaryKey;autoIncrement" json:"pricing_id"`
	Currency string `gorm:"size:3;not null;default:USD" json:"currency"`

	// unique -> 1:1 กับ Course
	CourseID uint    `gorm:"column:course_id;not null;uniqueIndex" json:"course_id"`
	Course   *Course `gorm:"foreignKey:CourseID;references:ID" json:"-"`

	DiscountPercent int        `gorm:"not null;default:0" json:"discount_percent"`
	PromoEndsAt     *time.Time `gorm:"column:promo_ends_at" json:"promo_ends_at,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (CoursePricing) TableName() string { return "course_pricings" }
