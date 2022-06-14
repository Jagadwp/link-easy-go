package repositories

import (
	"time"

	"gorm.io/gorm"

	"github.com/Jagadwp/link-easy-go/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) InsertUser(username, fullname, email, password string) (*models.User, error) {
	hashedPass, _ := models.Hash(password)

	var user *models.User = &models.User{
		Username:  username,
		Fullname:  fullname,
		Email:     email,
		Password:  string(hashedPass),
		Admin:     false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func (u *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
