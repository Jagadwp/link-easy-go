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
		return dto.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	response, err := ctr.services.InsertUrl(&req)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully inserted", response)
}

func (ctr *UrlController) GetAllUrlsByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Parameter id is not valid")
	}

	response, err := ctr.services.GetAllUrlsByUserID(userID)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}
	if len(*response) == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, "Url not found")
	}

	return dto.SuccessResponse(c, http.StatusOK, "Urls successfully fetched", response)
}

func (ctr *UrlController) GetUrlById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Parameter id is not valid")
	}

	response, err := ctr.services.GetUrlById(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}
	if (*response).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, "Url not found")
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully fetched", response)
}

func (ctr *UrlController) UpdateUrl(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Parameter id is not valid")
	}

	getUrlResponse, err := ctr.services.GetUrlById(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}
	if (*getUrlResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, "Url not found")
	}

	req := dto.UpdateUrlRequest{}
	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	updateResponse, err := ctr.services.UpdateUrl(id, &req)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully updated", updateResponse)
}

func (ctr *UrlController) DeleteUrl(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, "Parameter id is not valid")
	}

	getUrlResponse, err := ctr.services.GetUrlById(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}
	if (*getUrlResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, "Url not found")
	}

	deleteResponse, err := ctr.services.DeleteUrl(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully deleted", deleteResponse)
}
