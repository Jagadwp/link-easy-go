package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id" form:"id" gorm:"primaryKey"`
	Username  string    `json:"username" form:"username" gorm:"unique"`
	Fullname  string    `json:"fullname" form: "fullname"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Admin     bool      `json:"admin" form: "admin" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" form: "created_at"`
	UpdatedAt time.Time `json:"updated_at" form: "updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
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
