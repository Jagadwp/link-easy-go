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

	ErrRequiredFieldsNotValid = fmt.Errorf("required fields are empty or not valid")

	ErrJWTInvalid = fmt.Errorf("invalid token")
	ErrJWTExpired = fmt.Errorf("token already expired")

	
	ErrUrlShortLinkAlreadyExist = fmt.Errorf("short link already exist")
	ErrUrlNotFound = fmt.Errorf("url not found")
	ErrOriginalUrlNotValid = fmt.Errorf("original url is not valid")
	
	ErrUserAlreadyExist = fmt.Errorf("user already exist")
	ErrUserNotFound = fmt.Errorf("user not found")
	ErrUserWrongPass = fmt.Errorf("hash and password doesn't match")
	ErrFailedToProcessToken = fmt.Errorf("failed to process token")
)

type ErrorGorm struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}