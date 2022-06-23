package config

import "github.com/golang-jwt/jwt"

const (
	ENV_PATH          = ".env"
	APP_PORT          = "22347"
	SECRET_NANOID     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	JWT_SIGNATURE_KEY = "LOVEY-DOVEY-KEY"
)

var (
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)
