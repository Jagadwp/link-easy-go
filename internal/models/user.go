package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" form:"id" gorm:"primaryKey;not null"`
	Username  string    `json:"username" form:"username" gorm:"unique;not null"`
	Fullname  string    `json:"fullname" form: "fullname gorm:"not null"`
	Email     string    `json:"email" form:"email" gorm:"unique;not null"`
	Password  string    `json:"password" form:"password" gorm:"not null"`
	Admin     bool      `json:"admin" form: "admin" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" form: "created_at"`
	UpdatedAt time.Time `json:"updated_at" form: "updated_at"`
	Urls      []Url
}

//Habis implement validator
// type User struct {
// 	ID        int       `json:"id" form:"id" gorm:"primaryKey"`
// 	Username  string    `json:"username" form:"username" gorm:"unique" validate:"required,alphanumunicode"`
// 	Fullname  string    `json:"fullname" form: "fullname" validate:"required,alphanumunicode"`
// 	Email     string    `json:"email" form:"email" validate:"required,email"`
// 	Password  string    `json:"password" form:"password" validate:"required,gte=8,lte=32"`
// 	Admin     bool      `json:"admin" form: "admin" gorm:"default:false" validate:"boolean"`
// 	CreatedAt time.Time `json:"created_at" form: "created_at"`
// 	UpdatedAt time.Time `json:"updated_at" form: "updated_at"`
// }
