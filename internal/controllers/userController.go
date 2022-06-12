package controllers

import (
	"github.com/labstack/echo/v4"

	"github.com/Jagadwp/link-easy-go/internal/services"
)

const (
	Error = "error"
	Data  = "data"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (ctr *UserController) register(c echo.Context) {

	// req := data_object.RegisterResponse

	// response, err := ctr.service.Register(&req)

	// if err != nil {
	// 	c.String(http.StatusOK, "Hello, this is echo!")
	// }

	// c.JSON(http.StatusOK, response)

}
