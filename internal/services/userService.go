package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
	"github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/services/helper"
	"github.com/Jagadwp/link-easy-go/internal/shared"
	"github.com/Jagadwp/link-easy-go/internal/shared/config"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
	"github.com/dgrijalva/jwt-go"
	goJwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	usersRepo repositories.IUserRepository
}

func NewUserService(usersRepo repositories.IUserRepository) *UserService {
	return &UserService{usersRepo: usersRepo}
}

func (s *UserService) Login(req *dto.LoginRequest) (*dto.LoginResponse, int, error) {
	user, err := s.usersRepo.GetUserByUsername(req.Username)

	if err != nil {
		return &dto.LoginResponse{}, http.StatusBadRequest, shared.ErrUserNotFound
	}

	match, _ := helper.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &dto.LoginResponse{}, http.StatusBadRequest, shared.ErrUserWrongPass
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
		return &dto.LoginResponse{}, http.StatusInternalServerError, shared.ErrFailedToProcessToken
	}

	return &dto.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
		Admin:    user.Admin,
		Token:    signedToken,
	}, http.StatusOK, nil
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

func (s *UserService) InsertUser(req *dto.InsertUserRequest) (*dto.CommonUserResponse, int, error) {
	hashedPass, _ := helper.Hash(req.Password)

	user, err := s.usersRepo.InsertUser(&models.User{
		Username:  req.Username,
		Fullname:  req.Fullname,
		Email:     req.Email,
		Password:  hashedPass,
		Admin:     false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return &dto.CommonUserResponse{}, http.StatusInternalServerError, shared.ErrFailedToProcessRequest
	}

	return &dto.CommonUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Admin:     user.Admin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, http.StatusCreated, nil
}

func (s *UserService) GetAllUsers() (*[]dto.CommonUserResponse, int, error) {
	users, err := s.usersRepo.GetAllUsers()

	if err != nil {
		return &[]dto.CommonUserResponse{}, http.StatusInternalServerError, shared.ErrFailedToProcessRequest
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

	return &data, http.StatusOK, nil
}

func (s *UserService) GetUserById(id int) (*dto.CommonUserResponse, int, error) {
	user, err := s.usersRepo.GetUserById(id)

	if (*user).ID == 0 {
		return nil, http.StatusNotFound, shared.ErrUserNotFound
	}

	if err != nil {
		return nil, http.StatusInternalServerError, shared.ErrFailedToProcessRequest
	}

	return &dto.CommonUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Admin:     user.Admin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, http.StatusOK, nil

}

func (s *UserService) UpdateUser(id int, req *dto.UpdateUserRequest) (*dto.CommonUserResponse, int, error) {
	user, err := s.usersRepo.GetUserById(id)

	if err != nil {
		return &dto.CommonUserResponse{}, http.StatusNotFound, shared.ErrUserNotFound
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
		return &dto.CommonUserResponse{}, http.StatusInternalServerError, shared.ErrFailedToProcessRequest
	}

	return &dto.CommonUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Admin:     user.Admin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, http.StatusOK, nil
}

func (s *UserService) DeleteUser(id int) (*models.User, int, error) {
	user, err := s.usersRepo.GetUserById(id)

	if err != nil {
		return &models.User{}, http.StatusNotFound, shared.ErrUserNotFound
	}

	user, err = s.usersRepo.DeleteUser(user)

	if err != nil {
		return &models.User{}, http.StatusInternalServerError, shared.ErrFailedToProcessRequest
	}

	return user, http.StatusOK, nil
}
