package services

import (
	"time"

	"github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/services/helper"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
)

type UserService struct {
	usersRepo *repositories.UserRepository
}

func NewUserService(usersRepo *repositories.UserRepository) *UserService {
	return &UserService{usersRepo: usersRepo}
}

func (s *UserService) InsertUser(req *dto.InsertUserRequest) (*dto.CommonUserResponse, error) {
	hashedPass, _ := helper.Hash(req.Password)

	user, err := s.usersRepo.InsertUser(req.Username, req.Fullname, req.Email, string(hashedPass))

	if err != nil {
		return &dto.CommonUserResponse{}, err
	}

	return &dto.CommonUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Password:  user.Password,
		Admin:     user.Admin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) GetAllUsers() (*[]dto.CommonUserResponse, error) {
	users, err := s.usersRepo.GetAllUsers()

	if err != nil {
		return &[]dto.CommonUserResponse{}, err
	}

	var data []dto.CommonUserResponse

	for _, user := range *users {
		tempData := dto.CommonUserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Fullname:  user.Fullname,
			Email:     user.Email,
			Password:  user.Password,
			Admin:     user.Admin,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		data = append(data, tempData)
	}

	return &data, nil
}

func (s *UserService) GetUserById(id int) (*dto.CommonUserResponse, error) {
	user, err := s.usersRepo.GetUserById(id)

	if err != nil {
		return &dto.CommonUserResponse{}, err
	}

	return &dto.CommonUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Password:  user.Password,
		Admin:     user.Admin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}

func (s *UserService) UpdateUser(id int, req *dto.UpdateUserRequest) (*dto.CommonUserResponse, error) {
	user, err := s.usersRepo.GetUserById(id)

	if err != nil {
		return &dto.CommonUserResponse{}, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}

	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}

	if req.Password != "" {
		hashedPass, _ := helper.Hash(req.Password)
		user.Password = string(hashedPass)
	}

	user.Admin = req.Admin

	user.UpdatedAt = time.Now()

	user, err = s.usersRepo.UpdateUser(user)

	if err != nil {
		return &dto.CommonUserResponse{}, err
	}

	return &dto.CommonUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Password:  user.Password,
		Admin:     user.Admin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) DeleteUser(id int) error {
	user, err := s.usersRepo.GetUserById(id)

	if err != nil {
		return err
	}

	err = s.usersRepo.DeleteUser(user)

	if err != nil {
		return err
	}

	return nil
}
