package config

import "github.com/golang-jwt/jwt"

const (
	EnvPath = ".env"
	AppPort = "6000"

	JWT_SIGNATURE_KEY = "LOVEY-DOVEY-KEY"
)

var (
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)