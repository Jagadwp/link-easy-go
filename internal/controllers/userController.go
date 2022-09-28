package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/Jagadwp/link-easy-go/internal/services"
	"github.com/Jagadwp/link-easy-go/internal/shared"
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
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrRequiredFieldsNotValid.Error())
	}

	response, statusCode, err := ctr.services.Login(&req)

	if err != nil {
		return dto.ErrorResponse(c, statusCode, err.Error())
	}

	return dto.SuccessResponse(c, statusCode, "Login Success", response)
}

func (ctr *UserController) Register(c echo.Context) error {
	req := dto.InsertUserRequest{}

	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	response, statusCode, err := ctr.services.InsertUser(&req)

	if err != nil {
		return dto.ErrorResponse(c, statusCode, err.Error())
	}

	return dto.SuccessResponse(c, statusCode, "User successfully inserted", response)
}

func (ctr *UserController) GetAllUsers(c echo.Context) error {
	response, statusCode, err := ctr.services.GetAllUsers()

	if err != nil {
		return dto.ErrorResponse(c, statusCode, err.Error())
	}

	return dto.SuccessResponse(c, statusCode, "Users successfully fetched", response)

}

func (ctr *UserController) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrParamIDNotValid.Error())
	}

	response, statusCode, err := ctr.services.GetUserById(id)

	if err != nil {
		return dto.ErrorResponse(c, statusCode, err.Error())
	}

	return dto.SuccessResponse(c, statusCode, "User successfully fetched", response)
}

func (ctr *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	req := dto.UpdateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrRequiredFieldsNotValid.Error())
	}

	response, statusCode, err := ctr.services.UpdateUser(id, &req)

	if err != nil {
		return dto.ErrorResponse(c, statusCode, err.Error())
	}

	return dto.SuccessResponse(c, statusCode, "User successfully updated", response)
}

func (ctr *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrRequiredFieldsNotValid.Error())
	}

	deleteResponse, statusCode, err := ctr.services.DeleteUser(id)

	if err != nil {
		return dto.ErrorResponse(c, statusCode, err.Error())
	}

	return dto.SuccessResponse(c, statusCode, "User successfully deleted", deleteResponse)
}
