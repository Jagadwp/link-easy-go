package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/Jagadwp/link-easy-go/internal/services"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
)

type UserController struct {
	services *services.UserService
}

func NewUserController(services *services.UserService) *UserController {
	return &UserController{services: services}
}

func (ctr *UserController) Login(c echo.Context) error {
	req := dto.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	response, err := ctr.services.Login(&req)

	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return dto.SuccessResponse(c, http.StatusOK, "Login Success", response)
}

func (ctr *UserController) Register(c echo.Context) error {
	req := dto.InsertUserRequest{}

	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	response, err := ctr.services.InsertUser(&req)

	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "User successfully inserted", response)
}

func (ctr *UserController) GetAllUsers(c echo.Context) error {
	response, err := ctr.services.GetAllUsers()

	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "Users successfully fetched", response)

}

func (ctr *UserController) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Parameter id is not valid")
	}

	response, err := ctr.services.GetUserById(id)

	if (*response).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	return dto.SuccessResponse(c, http.StatusOK, "User successfully fetched", response)
}

func (ctr *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Parameter id is not valid")
	}

	getUserResponse, err := ctr.services.GetUserById(id)

	if (*getUserResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	req := dto.UpdateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	response, err := ctr.services.UpdateUser(id, &req)

	if err != nil {
		dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "User successfully updated", response)
}

func (ctr *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Parameter id is not valid")
	}

	getUserResponse, err := ctr.services.GetUserById(id)

	if (*getUserResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	deleteResponse, err := ctr.services.DeleteUser(id)

	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "User successfully deleted", deleteResponse)
}
