package controllers

import (
	"net/http"
	"strconv"

	"github.com/Jagadwp/link-easy-go/internal/services"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
	"github.com/labstack/echo/v4"
)
type UrlController struct {
  services *services.UrlService
}

func NewUrlController(services *services.UrlService) *UrlController {
	return &UrlController{services: services}
}

func (ctr *UrlController) InsertUrl(c echo.Context) error {
	req := dto.InsertUrlRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response, err := ctr.services.InsertUrl(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, response)
}

func (ctr *UrlController) GetAllUrlsByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response, err := ctr.services.GetAllUrlsByUserID(userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, response)
}