package services

import (
	"net/url"
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
	repository "github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/services/helper"
	"github.com/Jagadwp/link-easy-go/internal/shared"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
)

type UrlService struct {
	urlsRepo *repository.UrlRepository
}

func NewUrlService(urlsRepo *repository.UrlRepository) *UrlService {
	return &UrlService{urlsRepo: urlsRepo}
}

func (s *UrlService) GetAllUrlsByUserID(userID int) (*[]models.Url, error) {
	urls, err := s.urlsRepo.GetUrlsByUserID(userID)
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

func (s *UrlService) GetUrlByShortLink(shortLink string) (*dto.PublicUrlResponse, error) {
	url, err := s.urlsRepo.GetUrlByShortLink(shortLink)
	if err != nil {
		return &dto.PublicUrlResponse{}, err
	}

	return &dto.PublicUrlResponse{
		ShortLink: url.ShortLink,
		OriginalLink: url.OriginalLink,
	}, nil
}

func (s *UrlService) CreateUrl(req *dto.CreateUrlRequest) (*dto.CreateUrlResponse, error) {
	shortLink, errNanoId := helper.GenerateLink(); if errNanoId != nil {
		return &dto.CreateUrlResponse{}, errNanoId
	}

	if !s.IsUrlValid(req.OriginalLink) {
		return &dto.CreateUrlResponse{}, shared.ErrOriginalUrlNotValid
	}

	url, err := s.urlsRepo.CreateUrl(&models.Url{
		Title: req.Title,
		ShortLink: shortLink,
		OriginalLink: req.OriginalLink,
		HitCounter:   0,
		UserID:       req.UserID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}); if(err != nil) {
		return &dto.CreateUrlResponse{}, err
	}

	return &dto.CreateUrlResponse{
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

func (s *UrlService) UpdateUrl(id int, req *dto.UpdateUrlRequest) (*dto.UpdateUrlResponse, error) {
	url, _ := s.urlsRepo.GetUrlById(id)
	if (*url).ID == 0 {
		return &dto.UpdateUrlResponse{}, shared.ErrUrlNotFound
	}

	url.Title = req.Title
	url.ShortLink = req.ShortLink
	url.OriginalLink = req.OriginalLink
	url.UpdatedAt = time.Now()

	urlByShortLink, _ := s.urlsRepo.GetUrlByShortLink(req.ShortLink)
	if ((*urlByShortLink).ID != 0 && (*urlByShortLink).ID != id) {
		return &dto.UpdateUrlResponse{}, shared.ErrUrlShortLinkAlreadyExist
	}

	if !s.IsUrlValid(req.OriginalLink) {
		return &dto.UpdateUrlResponse{}, shared.ErrOriginalUrlNotValid
	}

	if !s.IsUserAllowedToEdit(req.UserID, url.UserID) {
		return &dto.UpdateUrlResponse{}, shared.ErrForbiddenToAccess
	}

	url, err := s.urlsRepo.UpdateUrl(url)
	if err != nil {
		return &dto.UpdateUrlResponse{}, err
	}

	return &dto.UpdateUrlResponse{
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

func (s *UrlService) DeleteUrl(id int, userID int) (*dto.UpdateUrlResponse, error) {
	url, _ := s.urlsRepo.GetUrlById(id)
	if (*url).ID == 0 {
		return &dto.UpdateUrlResponse{}, shared.ErrUrlNotFound
	}

	if !s.IsUserAllowedToEdit(userID, url.UserID) {
		return &dto.UpdateUrlResponse{}, shared.ErrForbiddenToAccess
	}

	url, err := s.urlsRepo.DeleteUrl(url)
	if err != nil {
		return &dto.UpdateUrlResponse{}, err
	}

	return &dto.UpdateUrlResponse{
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
