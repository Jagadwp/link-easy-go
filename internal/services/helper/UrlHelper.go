package helper

import (
	"net/url"

	"github.com/Jagadwp/link-easy-go/internal/shared/config"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateLink() (string, error) {
	shortLink, errNanoId := nanoid.Generate(config.SECRET_NANOID, 6)

	if errNanoId != nil {
		return "", errNanoId
	}

	return shortLink, nil
}

func IsUserAllowedToEditUrl(userID int, userIDInUrl int) (bool) {
	return userID == userIDInUrl
}

func IsUrlValid(toTest string) (bool) {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}