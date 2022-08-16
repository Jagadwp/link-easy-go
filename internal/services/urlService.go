package services

import (
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
	"github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/services/helper"
	"github.com/Jagadwp/link-easy-go/internal/shared"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
)

var IsUserAllowedToEditUrl = helper.IsUserAllowedToEditUrl
var IsUrlValid = helper.IsUrlValid
var GenerateLink = helper.GenerateLink
var Now = time.Now

type UrlService struct {
	urlsRepo repositories.IUrlRepository
}

func NewUrlService(urlsRepo repositories.IUrlRepository) *UrlService {
	return &UrlService{urlsRepo: urlsRepo}
}

// public
func (s *UrlService) GetUrlPublicByShortLink(shortLink string) (*dto.PublicUrlResponse, error) {
	url, _ := s.urlsRepo.GetUrlByShortLink(shortLink)
	if (*url).ID == 0 {
		return &dto.PublicUrlResponse{}, shared.ErrUrlNotFound
	}

	s.urlsRepo.IncrementHitCounter(url)

	return &dto.PublicUrlResponse{
		ShortLink: url.ShortLink,
		OriginalLink: url.OriginalLink,
	}, nil
}

func (s *UrlService) GetUrlsByUserID(userID int) (*[]models.Url, error) {
	urls, _ := s.urlsRepo.GetUrlsByUserID(userID)
	if len(*urls) == 0 {
		return &[]models.Url{}, shared.ErrUrlNotFound
	}

	return urls, nil
}

func (s *UrlService) GetUrlUserById(id int, userID int) (*models.Url, error) {
	url, _ := s.urlsRepo.GetUrlById(id)
	if (*url).ID == 0 {
		return &models.Url{}, shared.ErrUrlNotFound
	}

	if !IsUserAllowedToEditUrl(userID, url.UserID) {
		return &models.Url{}, shared.ErrForbiddenToAccess
	}

	return url, nil
}

func (s *UrlService) CreateUrl(req *dto.CreateUrlRequest) (*dto.CreateUrlResponse, error) {
	shortLink, errNanoId := GenerateLink()
	if errNanoId != nil {
		return &dto.CreateUrlResponse{}, errNanoId
	}

	if !IsUrlValid(req.OriginalLink) {
		return &dto.CreateUrlResponse{}, shared.ErrOriginalUrlNotValid
	}

	now := Now()
	url, err := s.urlsRepo.CreateUrl(&models.Url{
		Title: req.Title,
		ShortLink: shortLink,
		OriginalLink: req.OriginalLink,
		HitCounter:   0,
		UserID:       req.UserID,
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if(err != nil) {
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
	url.UpdatedAt = Now()

	urlByShortLink, _ := s.urlsRepo.GetUrlByShortLink(req.ShortLink)
	if ((*urlByShortLink).ID != 0 && (*urlByShortLink).ID != id) {
		return &dto.UpdateUrlResponse{}, shared.ErrUrlShortLinkAlreadyExist
	}

	if !IsUrlValid(req.OriginalLink) {
		return &dto.UpdateUrlResponse{}, shared.ErrOriginalUrlNotValid
	}

	if !IsUserAllowedToEditUrl(req.UserID, url.UserID) {
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

	if !IsUserAllowedToEditUrl(userID, url.UserID) {
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


