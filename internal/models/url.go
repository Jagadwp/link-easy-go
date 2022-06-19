package models

import (
	"time"
)

type Url struct {
	ID           int       `json:"id" form:"id" gorm:"primaryKey;not null"`
	Title        string    `json:"title" form:"title"`
	ShortLink    string    `json:"short_link" form:"short_link" gorm:"unique;not null"`
	OriginalLink string    `json:"original_link" form:"original_link" gorm:"not null"`
	HitCounter   int       `json:"hit_counter" form:"hit_counter"`
	UserID       *int       `json:"user_id" form:"user_id"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}
