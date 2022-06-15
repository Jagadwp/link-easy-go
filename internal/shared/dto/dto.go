package dto

import "time"

type InsertUserRequest struct {
	Username  string    `json:"username" form:"username"`
	Fullname  string    `json:"fullname" form: "fullname"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
}

type CommonUserResponse struct {
	ID        int       `json:"id" form:"id"`
	Username  string    `json:"username" form:"username"`
	Fullname  string    `json:"fullname" form: "fullname"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Admin     bool      `json:"admin" form:"admin""`
	CreatedAt time.Time `json:"created_at" form: "created_at"`
	UpdatedAt time.Time `json:"updated_at" form: "updated_at"`
}

type InsertUrlRequest struct {
	Title        string    `json:"title" form:"title"`
	ShortLink    string    `json:"short_link" form:"short_link"`
	OriginalLink string    `json:"original_link" form:"original_link"`
	UserID       int       `json:"user_id" form:"user_id"`
}

type InsertUrlResponse struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	ShortLink     string    `json:"short_link"`
	OriginalLink  string    `json:"original_link"`
	HitCounter    int       `json:"hit_counter"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdateUrlRequest struct {
	ID           int       `json:"id"`
	Title        string    `json:"title" form:"title"`
	ShortLink    string    `json:"short_link" form:"short_link"`
	OriginalLink string    `json:"original_link" form:"original_link"`
}

type UpdateUrlResponse struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	ShortLink     string    `json:"short_link"`
	OriginalLink  string    `json:"original_link"`
	HitCounter    int       `json:"hit_counter"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
