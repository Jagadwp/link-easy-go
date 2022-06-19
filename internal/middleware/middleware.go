package middleware

import (
	"github.com/Jagadwp/link-easy-go/internal/shared/config"
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(config.JWT_SECRET),
})
