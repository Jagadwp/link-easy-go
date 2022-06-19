package helper

import (
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
