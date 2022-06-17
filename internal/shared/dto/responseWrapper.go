package dto

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Success  bool `json:"success"`
	Message string `json:"message"`
	ErrorCode  int    `json:"error_code"`
	Data interface{} `json:"data"`
}

type Success struct {
	Success  bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return Response(c, statusCode, Success{
		Success: true,
		Message: message,
		Data: data,
	})
}

func ErrorResponse(c echo.Context, statusCode int, err string) error {
	return Response(c, statusCode, Error{
		Success: false,
		Message: err,
		ErrorCode: statusCode,
		Data: nil,
	})
}
