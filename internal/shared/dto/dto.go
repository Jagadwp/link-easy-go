package dto

import "time"

type RegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type InsertUrlRequest struct {
	ShortLink string `json:"short_link" form:"short_link"`
	OriginalLink string `json:"original_link" form:"original_link"`
	UserID int `json:"user_id" form:"user_id"`
}

type InsertUrlResponse struct {
	ID       int64  `json:"id"`
	ShortLink string `json:"short_link"`
	OriginalLink string `json:"original_link"`
	HitCounter int `json:"hit_counter"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type ErrorResponse struct {
// 	Message string `json:"message"`
// }
