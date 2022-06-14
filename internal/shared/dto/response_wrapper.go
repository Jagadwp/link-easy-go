package dto

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponseCode string

//List of success response status
const (
	Success SuccessResponseCode = "success"
)

//SuccessResponse default payload response
type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

//NewSuccessResponse create new success payload
func NewSuccessResponse(data interface{}) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		"Success",
		data,
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponseWithoutData() (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		"Success",
		map[string]interface{}{},
	}
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return c.JSON(statusCode, data)
}

func MessageResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
