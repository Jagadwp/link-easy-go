package controllers

import (
	"net/http"

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

// func (ctr *UserController) test(c echo.Context) error {
	
// 	db.Find(&user)
// 	return c.JSON(200, response)

// }
