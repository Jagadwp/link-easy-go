package shared

import "fmt"

const (
	CODE_ERROR_DUPLICATE_KEY = "23505"
)

var (
	ErrBadRequest          = fmt.Errorf("bad request")
	ErrInternalServerError = fmt.Errorf("internal server error")
	ErrFailedToProcessRequest = fmt.Errorf("failed to process request")
	ErrForbiddenToAccess = fmt.Errorf("you are forbidden to access this resource")
	ErrParamIDNotValid = fmt.Errorf("parameter id is not valid")

	ErrJWTInvalid = fmt.Errorf("invalid token")
	ErrJWTExpired = fmt.Errorf("token already expired")

	ErrUserAlreadyExist = fmt.Errorf("user already exist")

	ErrUrlShortLinkAlreadyExist = fmt.Errorf("short link already exist")
	ErrUrlNotFound = fmt.Errorf("url not found")
)

type ErrorGorm struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}