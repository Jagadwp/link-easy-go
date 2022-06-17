package models

import (
	"time"
)

type Url struct {
	ID           int       `json:"id" form:"id" gorm:"primaryKey;not null"`
	Title        string    `json:"title" form:"title" gorm:"not null"`
	ShortLink    string    `json:"short_link" form:"short_link" gorm:"unique;not null"`
	OriginalLink string    `json:"original_link" form:"original_link" gorm:"not null"`
	HitCounter   int       `json:"hit_counter" form:"hit_counter"`
	CreatedBy    int       `json:"created_by" form:"created_by"`
	CreatedAt    time.Time `json:"created_at" form: "created_at"`
	UpdatedAt    time.Time `json:"updated_at" form: "updated_at"`
	UserID       int
}
