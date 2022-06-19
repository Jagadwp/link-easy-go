package services

import (
	"net/url"
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
	repository "github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/services/helper"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
)

type UrlService struct {
	urlsRepo *repository.UrlRepository
}

func NewUrlService(urlsRepo *repository.UrlRepository) *UrlService {
	return &UrlService{urlsRepo: urlsRepo}
}

func (s *UrlService) CreateShortUrl(req *dto.GenerateUrlRequest) (*dto.InsertUrlResponse, error) {
	shortLink, errNanoId := helper.GenerateLink()
	shortLink = "https://linkeasy.in/" + shortLink

	if errNanoId != nil {
		return &dto.InsertUrlResponse{}, errNanoId
	}

	url, err := s.urlsRepo.CreateShortUrl(req.Title, req.OriginalLink, shortLink, req.UserID)

	if err != nil {
		return &dto.InsertUrlResponse{}, err
	}

	return &dto.InsertUrlResponse{
		ID:           url.ID,
		Title:        url.Title,
		ShortLink:    url.ShortLink,
		OriginalLink: url.OriginalLink,
		HitCounter:   url.HitCounter,
		UserID:       url.UserID,
		CreatedAt:    url.CreatedAt,
		UpdatedAt:    url.UpdatedAt,
	}, nil
}

func (s *UrlService) InsertUrl(title, shortLink, originalLink string, id *int) (*dto.InsertUrlResponse, error) {
	shortLink = "https://linkeasy.in/" + shortLink
	
	url, err := s.urlsRepo.InsertUrl(title, shortLink, originalLink, id)

	if err != nil {
		return &dto.InsertUrlResponse{}, err
	}

	return &dto.InsertUrlResponse{
		ID:           url.ID,
		Title:        url.Title,
		ShortLink:    url.ShortLink,
		OriginalLink: url.OriginalLink,
		HitCounter:   url.HitCounter,
		UserID:       url.UserID,
		CreatedAt:    url.CreatedAt,
		UpdatedAt:    url.UpdatedAt,
	}, nil
}

func (s *UrlService) GetAllUrlsByUserID(userID int) (*[]models.Url, error) {
	urls, err := s.urlsRepo.GetAllUrlsByUserID(userID)
	if err != nil {
		return &[]models.Url{}, err
	}

	return urls, nil
}

func (s *UrlService) GetUrlById(id int) (*models.Url, error) {
	url, err := s.urlsRepo.GetUrlById(id)
	if err != nil {
		return &models.Url{}, err
	}

	return url, nil
}

func (s *UrlService) UpdateUrl(id int, req *dto.UpdateUrlRequest) (*dto.UpdateUrlResponse, error) {
	url, err := s.urlsRepo.GetUrlById(id)
	if err != nil {
		return &dto.UpdateUrlResponse{}, err
	}

	if req.Title != "" {
		url.Title = req.Title
	}
	if req.ShortLink != "" {
		url.ShortLink = req.ShortLink
	}
	if req.OriginalLink != "" {
		url.OriginalLink = req.OriginalLink
	}
	url.UpdatedAt = time.Now()

	url, err = s.urlsRepo.UpdateUrl(url)
	if err != nil {
		return &dto.UpdateUrlResponse{}, err
	}

	return &dto.UpdateUrlResponse{
		ID:           url.ID,
		Title:        url.Title,
		ShortLink:    url.ShortLink,
		OriginalLink: url.OriginalLink,
		HitCounter:   url.HitCounter,
		UserID:       *url.UserID,
		CreatedAt:    url.CreatedAt,
		UpdatedAt:    url.UpdatedAt,
	}, nil
}

func (s *UrlService) DeleteUrl(id int) (*models.Url, error) {
	url, err := s.urlsRepo.GetUrlById(id)
	if err != nil {
		return &models.Url{}, err
	}

	url, err = s.urlsRepo.DeleteUrl(url)
	if err != nil {
		return &models.Url{}, err
	}

	return url, nil
}

func (s *UrlService) IsUserAllowedToEdit(userID int, userIDInUrl int) (bool) {
	return userID == userIDInUrl
}

func (s *UrlService) IsUrlValid(toTest string) (bool) {
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
