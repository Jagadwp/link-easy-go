package services

import (
	repository "github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
)

type UserService struct {
	usersRepo *repository.UserRepository
}

func NewUserService(usersRepo *repository.UserRepository) *UserService {
	return &UserService{usersRepo: usersRepo}
}

func (s *UserService) InsertUser(req *dto.InsertUserRequest) (*dto.CommonUserResponse, error) {
	user, err := s.usersRepo.InsertUser(req.Username, req.Fullname, req.Email, req.Password)

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

func (s *UserService) GetAllUsers() ([]dto.CommonUserResponse, error) {
	users, err := s.usersRepo.GetAllUsers()

	if err != nil {
		return []dto.CommonUserResponse{}, err
	}

	var data []dto.CommonUserResponse

	for _, user := range users {
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

	return data, nil
}
