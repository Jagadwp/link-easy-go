package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Jagadwp/link-easy-go/internal/services"
	"github.com/Jagadwp/link-easy-go/internal/shared"
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
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.MESSAGE_FIELD_REQUIRED)
	}

	response, err := ctr.services.InsertUrl(&req)
	if err != nil {
		if errors.Is(err, shared.ErrUrlShortLinkAlreadyExist) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlShortLinkAlreadyExist.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
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
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if len(*response) == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
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
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if (*response).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
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
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if (*getUrlResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
	}

	req := dto.UpdateUrlRequest{}
	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.MESSAGE_FIELD_REQUIRED)
	}

	updateResponse, err := ctr.services.UpdateUrl(id, &req)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
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
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if (*getUrlResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
	}

	deleteResponse, err := ctr.services.DeleteUrl(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully deleted", deleteResponse)
}
