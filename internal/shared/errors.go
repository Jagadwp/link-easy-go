package shared

import "fmt"

var (
	ErrBadRequest          = fmt.Errorf("bad request")
	ErrInternalServerError = fmt.Errorf("internal server error")

	ErrUserAlreadyExist = fmt.Errorf("user already exist")
)
