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

func (s *UserService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// user, err := s.usersRepo.InsertUser(
	_, err := s.usersRepo.InsertUser(
		req.Username, req.Password,
	)

	if err != nil {
		return &dto.RegisterResponse{}, err
	}

	return &dto.RegisterResponse{
		Username: "jaahjasa",
	}, nil
}
