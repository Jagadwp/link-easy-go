package services

import (
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
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
	url, err := s.urlsRepo.InsertUrl(req.Title, req.ShortLink, req.OriginalLink, req.UserID)

	if(err != nil) {
		return &dto.InsertUrlResponse{}, err
	}

	return &dto.InsertUrlResponse{
		ID: url.ID,
		Title: url.Title,
		ShortLink: url.ShortLink,
		OriginalLink: url.OriginalLink,
		HitCounter: url.HitCounter,
		CreatedBy: url.CreatedBy,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}, nil
}

func (s *UrlService) GetAllUrlsByUserID(userID int) (*[]models.Url, error) {
	urls, err := s.urlsRepo.GetAllUrlsByUserID(userID)

	if(err != nil) {
		return &[]models.Url{}, err
	}

	return urls, nil
}

func (s *UrlService) UpdateUrl(id int, req *dto.UpdatetUrlRequest) (*dto.UpdateUrlResponse, error) {
	url, err := s.urlsRepo.GetUrlById(id)

	if(err != nil) {
		return &dto.UpdateUrlResponse{}, err
	}

	if (req.Title != "") {
		url.Title = req.Title
	}
	if (req.ShortLink != "") {
		url.ShortLink = req.ShortLink
	}
	if (req.OriginalLink != "") {
		url.OriginalLink = req.OriginalLink
	}
	url.UpdatedAt = time.Now()

	url, err = s.urlsRepo.UpdateUrl(url)

	if(err != nil) {
		return &dto.UpdateUrlResponse{}, err
	}

	return &dto.UpdateUrlResponse{
		ID: url.ID,
		Title: url.Title,
		ShortLink: url.ShortLink,
		OriginalLink: url.OriginalLink,
		HitCounter: url.HitCounter,
		CreatedBy: url.CreatedBy,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}, nil

}