package repositories

import (
	"gorm.io/gorm"

	"github.com/Jagadwp/link-easy-go/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetAllUsers(id int64) ([]*models.User, error) {
	var users []*models.User

	return users, nil

}

func (u *UserRepository) InsertUser(username, password string) (*models.User, error) {

	return &models.User{
		Username: username,
		Password: password,
	}, nil
}
