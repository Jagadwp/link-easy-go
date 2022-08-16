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
	urlService  *services.UrlService
	userService *services.UserService
}

func NewUrlController(
	urlService *services.UrlService,
	userService *services.UserService) *UrlController {
	return &UrlController{
		urlService:  urlService,
		userService: userService,
	}
}

// public
func (ctr *UrlController) GetUrlPublicByShortLink(c echo.Context) error {
	shortLink := c.Param("short_link")

	response, err := ctr.urlService.GetUrlPublicByShortLink(shortLink)
	if err != nil {
		if(errors.Is(err, shared.ErrUrlNotFound)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlNotFound.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully fetched", response)
}

// public
func (ctr *UrlController) RedirectShortLink(c echo.Context) error {
	shortLink := c.Param("short_link")

	response, err := ctr.urlService.GetUrlPublicByShortLink(shortLink)
	if err != nil {
		if(errors.Is(err, shared.ErrUrlNotFound)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlNotFound.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
	}

	return c.Redirect(http.StatusFound, response.OriginalLink)
}

func (ctr *UrlController) GetUrlsByUserID(c echo.Context) error {
	currentUser, ok := ctr.userService.GetCurrentUser(c)
	if !ok {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrJWTInvalid.Error())
	}

	userID := currentUser.ID

	response, err := ctr.urlService.GetUrlsByUserID(userID)
	if err != nil {
		if(errors.Is(err, shared.ErrUrlNotFound)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlNotFound.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
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

	response, err := ctr.urlService.GetUrlUserById(id, userID)
	if err != nil {
		if(errors.Is(err, shared.ErrUrlNotFound)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlNotFound.Error())
		} else if (errors.Is(err, shared.ErrForbiddenToAccess)) {
			return dto.ErrorResponse(c, http.StatusForbidden, shared.ErrForbiddenToAccess.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully fetched", response)
}

func (ctr *UrlController) CreateUrl(c echo.Context) error {
	currentUser, ok := ctr.userService.GetCurrentUser(c)
	if !ok {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrJWTInvalid.Error())
	}

	userID := currentUser.ID
	req := dto.CreateUrlRequest{}
	req.UserID = userID
	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrRequiredFieldsNotValid.Error())
	}
		
	response, err := ctr.urlService.CreateUrl(&req)
	if err != nil {
		if(errors.Is(err, shared.ErrOriginalUrlNotValid)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrOriginalUrlNotValid.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully created", response)
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
	req := dto.UpdateUrlRequest{}
	req.UserID = userID
	if err := c.Bind(&req); err != nil {
		return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrRequiredFieldsNotValid.Error())
	}

	response, err := ctr.urlService.UpdateUrl(id, &req)
	if err != nil {
		if(errors.Is(err, shared.ErrUrlNotFound)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlNotFound.Error())
		} else if(errors.Is(err, shared.ErrUrlShortLinkAlreadyExist)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlShortLinkAlreadyExist.Error())
		} else if(errors.Is(err, shared.ErrOriginalUrlNotValid)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrOriginalUrlNotValid.Error())
		} else if (errors.Is(err, shared.ErrForbiddenToAccess)) {
			return dto.ErrorResponse(c, http.StatusForbidden, shared.ErrForbiddenToAccess.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully updated", response)
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

	response, err := ctr.urlService.DeleteUrl(id, userID)
	if err != nil {
		if(errors.Is(err, shared.ErrUrlNotFound)) {
			return dto.ErrorResponse(c, http.StatusBadRequest, shared.ErrUrlNotFound.Error())
		} else if (errors.Is(err, shared.ErrForbiddenToAccess)) {
			return dto.ErrorResponse(c, http.StatusForbidden, shared.ErrForbiddenToAccess.Error())
		} else {
			return dto.ErrorResponse(c, http.StatusInternalServerError, shared.ErrFailedToProcessRequest.Error())
		}
	}

	return dto.SuccessResponse(c, http.StatusOK, "Url successfully deleted", response)
}
