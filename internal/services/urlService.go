package services

import (
	repository "github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
)

type UrlService struct {
	urlsRepo *repository.UrlRepository
}

func NewUrlService(urlsRepo *repository.UrlRepository) *UrlService {
	return &UrlService{urlsRepo: urlsRepo}
}

func (s *UrlService) InsertUrl(req *dto.InsertUrlRequest) (*dto.InsertUrlResponse, error) {
	url, err := s.urlsRepo.InsertUrl(req.ShortLink, req.OriginalLink, req.UserID)

	if(err != nil) {
		return &dto.InsertUrlResponse{}, err
	}

	return &dto.InsertUrlResponse{
		ShortLink: url.ShortLink,
		OriginalLink: url.OriginalLink,
		HitCounter: url.HitCounter,
		CreatedBy: url.CreatedBy,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}, nil
}
