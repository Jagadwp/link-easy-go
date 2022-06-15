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

func (ctr *UserController) InsertUser(c echo.Context) error {
	req := dto.InsertUserRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response, err := ctr.services.InsertUser(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(201, response)
}

func (ctr *UserController) GetAllUsers(c echo.Context) error {
	response, err := ctr.services.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, response)

}

func (ctr *UserController) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response, err := ctr.services.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, response)
}

func (ctr *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	req := dto.UpdateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response, err := ctr.services.UpdateUser(id, &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, response)
}

func (ctr *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := ctr.services.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, "User successfully deleted")
}
