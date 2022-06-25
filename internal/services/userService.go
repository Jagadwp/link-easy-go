package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
	"github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/services/helper"
	"github.com/Jagadwp/link-easy-go/internal/shared/config"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
	"github.com/dgrijalva/jwt-go"
	goJwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	usersRepo *repositories.UserRepository
}

func NewUserService(usersRepo *repositories.UserRepository) *UserService {
	return &UserService{usersRepo: usersRepo}
}

func (s *UserService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.usersRepo.GetUserByUsername(req.Username)

	if err != nil {
		return &dto.LoginResponse{}, errors.New("username not found")
	}

	match, _ := helper.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &dto.LoginResponse{}, errors.New("hash and password doesn't match")
	}

	claims := &dto.JwtCustomClaims{
		ID:       user.ID,
		Username: user.Username,
		Admin:    user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(config.JWT_SIGNING_METHOD, claims)

	fmt.Println(claims)

	signedToken, err := token.SignedString([]byte(config.JWT_SIGNATURE_KEY))
	if err != nil {
		return &dto.LoginResponse{}, errors.New("error while processing token")
	}

	return &dto.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
		Admin:    user.Admin,
		Token:    signedToken,
	}, nil
}

func (s *UserService) GetCurrentUser(c echo.Context) (userInfo *dto.JwtUserInfo, ok bool) {
	user, ok := c.Get("user").(*goJwt.Token)

	if !ok {
		return &dto.JwtUserInfo{}, ok
	}

	claims := user.Claims.(goJwt.MapClaims)
	username := claims["username"].(string)
	userID := claims["id"].(float64)
	isAdmin := claims["admin"].(bool)

	newUser := &dto.JwtUserInfo{
		ID:       int(userID),
		Username: username,
		Admin:    isAdmin,
	}

	return newUser, true
}

func (s *UserService) InsertUser(req *dto.InsertUserRequest) (*dto.CommonUserResponse, error) {
	hashedPass, _ := helper.Hash(req.Password)

	user, err := s.usersRepo.InsertUser(req.Username, req.Fullname, req.Email, hashedPass)

	if err != nil {
		return &dto.CommonUserResponse{}, err
	}

	return &dto.CommonUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
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

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		hashedPass, _ := helper.Hash(req.Password)
		user.Password = hashedPass
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
		Admin:     user.Admin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) DeleteUser(id int) (*models.User, error) {
	user, err := s.usersRepo.GetUserById(id)

	if err != nil {
		return &models.User{}, err
	}

	user, err = s.usersRepo.DeleteUser(user)

	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}
