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

	var user *models.User = &models.User{
		Username:  username,
		Fullname:  fullname,
		Email:     email,
		Password:  password,
		Admin:     false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User

	if err := u.db.First(&user, id).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	query := u.db.Where("username = ?", username).First(&user)

	if err := query.Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (u *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := u.db.Save(user).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (u *UserRepository) DeleteUser(user *models.User) (*models.User, error) {
	if err := u.db.Delete(user).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}
