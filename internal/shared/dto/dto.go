package dto

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

type JwtUserInfo struct {
	ID       int
	Username string
	Admin    bool
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
	Fullname string `json:"fullname" form: "fullname"`
	Email    string `json:"email" form:"email"`
	Admin    bool   `json:"admin" form:"admin""`
}

// User Data Transfer Object

type UpdateUserRequest struct {
	Username string `json:"username" form:"username"`
	Fullname string `json:"fullname" form: "fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Admin    bool   `json:"admin" form:"admin""`
}

type InsertUserRequest struct {
	Username string `json:"username" form:"username"`
	Fullname string `json:"fullname" form: "fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}


type CommonUserResponse struct {
	ID        int       `json:"id" form:"id"`
	Username  string    `json:"username" form:"username"`
	Fullname  string    `json:"fullname" form: "fullname"`
	Email     string    `json:"email" form:"email"`
	Admin     bool      `json:"admin" form:"admin""`
	CreatedAt time.Time `json:"created_at" form: "created_at"`
	UpdatedAt time.Time `json:"updated_at" form: "updated_at"`
}

// URL Data Transfer Object

type GenerateUrlRequest struct {
	Title        string `json:"title" form:"title"`
	OriginalLink string `json:"original_link" form:"original_link"`
	UserID       *int    `json:"user_id" form:"user_id"`
}

type InsertUrlRequest struct {
	Title        string `json:"title" form:"title"`
	ShortLink    string `json:"short_link" form:"short_link"`
	OriginalLink string `json:"original_link" form:"original_link"`
	UserID       *int    `json:"user_id" form:"user_id"`
}


type InsertUrlResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	ShortLink    string    `json:"short_link"`
	OriginalLink string    `json:"original_link"`
	HitCounter   int       `json:"hit_counter"`
	UserID       *int       `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateUrlRequest struct {
	Title        string `json:"title" form:"title"`
	ShortLink    string `json:"short_link" form:"short_link"`
	OriginalLink string `json:"original_link" form:"original_link"`
}

type UpdateUrlResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	ShortLink    string    `json:"short_link"`
	OriginalLink string    `json:"original_link"`
	HitCounter   int       `json:"hit_counter"`
	UserID       int       `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PublicUrlResponse struct {
	ShortLink    string    `json:"short_link"`
	OriginalLink string    `json:"original_link"`
}
