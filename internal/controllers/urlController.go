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
	urlService *services.UrlService
	userService *services.UserService
}

func NewUrlController(
	urlService *services.UrlService,
	userService *services.UserService) *UrlController {
	return &UrlController{
		urlService: urlService,
		userService: userService,
	}
}

func (ctr *UrlController) CreateShortUrl(c echo.Context) error {
	req := dto.GenerateUrlRequest{}

	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.MESSAGE_FIELD_REQUIRED)
	}

	response, err := ctr.urlService.CreateShortUrl(&req)

	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, "Failed to process request")
	}

	return dto.SuccessResponse(c, http.StatusOK, "Generated Url successfully inserted", response)
}

func (ctr *UrlController) InsertUrl(c echo.Context) error {
	currentUser, ok := ctr.userService.GetCurrentUser(c)
	if !ok {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrJWTInvalid.Error())
	}

	userID := currentUser.ID
	req := dto.InsertUrlRequest{}
	
	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.MESSAGE_FIELD_REQUIRED)
	}
	if !ctr.urlService.IsUrlValid(req.OriginalLink) {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrOriginalUrlNotValid.Error())
	}

	response, err := ctr.urlService.InsertUrl(req.Title, req.ShortLink, req.OriginalLink, &userID)
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
	currentUser, ok := ctr.userService.GetCurrentUser(c)
	if !ok {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrJWTInvalid.Error())
	}

	userID := currentUser.ID

	response, err := ctr.urlService.GetAllUrlsByUserID(userID)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if len(*response) == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
	}

	return dto.SuccessResponse(c, http.StatusOK, "Urls successfully fetched", response)
}

func (ctr *UrlController) GetUrlUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrParamIDNotValid.Error())
	}

	currentUser, ok := ctr.userService.GetCurrentUser(c)
	if !ok {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrJWTInvalid.Error())
	}

	userID := currentUser.ID

	response, err := ctr.urlService.GetUrlById(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if (*response).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
	}

	if (ctr.urlService.IsUserAllowedToEdit(userID, *response.UserID)) {
		return dto.SuccessResponse(c, http.StatusOK, "Url successfully fetched", response)
	} else {
		return dto.ErrorResponse(c, http.StatusForbidden, shared.ErrForbiddenToAccess.Error())
	}
}

func (ctr *UrlController) UpdateUrl(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrParamIDNotValid.Error())
	}

	currentUser, ok := ctr.userService.GetCurrentUser(c)
	if !ok {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrJWTInvalid.Error())
	}

	userID := currentUser.ID

	getUrlResponse, err := ctr.urlService.GetUrlById(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if (*getUrlResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
	}

	if (ctr.urlService.IsUserAllowedToEdit(userID, *getUrlResponse.UserID)) {
		req := dto.UpdateUrlRequest{}
		if err := c.Bind(&req); err != nil {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.MESSAGE_FIELD_REQUIRED)
		}

		updateResponse, err := ctr.urlService.UpdateUrl(id, &req)
		if err != nil {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}

		return dto.SuccessResponse(c, http.StatusOK, "Url successfully updated", updateResponse)	
	} else {
		return dto.ErrorResponse(c, http.StatusForbidden, shared.ErrForbiddenToAccess.Error())
	}
}

func (ctr *UrlController) DeleteUrl(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrParamIDNotValid.Error())
	}

	currentUser, ok := ctr.userService.GetCurrentUser(c)
	if !ok {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrJWTInvalid.Error())
	}

	userID := currentUser.ID

	getUrlResponse, err := ctr.urlService.GetUrlById(id)
	if err != nil {
		return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
	}
	if (*getUrlResponse).ID == 0 {
		return dto.ErrorResponse(c, http.StatusNotFound, shared.ErrUrlNotFound.Error())
	}

	if (ctr.urlService.IsUserAllowedToEdit(userID, *getUrlResponse.UserID)) {
		deleteResponse, err := ctr.urlService.DeleteUrl(id)
		if err != nil {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}

		return dto.SuccessResponse(c, http.StatusOK, "Url successfully deleted", deleteResponse)
	} else {
		return dto.ErrorResponse(c, http.StatusForbidden, shared.ErrForbiddenToAccess.Error())
	}

	
}
