package tests

import (
	"testing"
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
	"github.com/Jagadwp/link-easy-go/internal/services"
	"github.com/Jagadwp/link-easy-go/internal/shared"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
	"github.com/Jagadwp/link-easy-go/mocks"
	"gopkg.in/go-playground/assert.v1"
)

// var IsUserAllowedToEditUrl = services.IsUserAllowedToEditUrl
// var IsUrlValid = services.IsUrlValid
// var NewUrlService = services.NewUrlService
// var Now = services.Now
// var GenerateLink = helper.GenerateLink

func TestGetUrlUserById(t *testing.T) {
	oldUrl := &models.Url{
		ID: 1,
		Title: "Test Title",
		ShortLink: "generatedLink",
		OriginalLink: "http://localhost:8080/test",
		HitCounter:   0,
		UserID:       1,
		CreatedAt:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	cases := []struct{
		name string
		IsUserAllowedToEditUrl func(userID int, userIDInUrl int) bool
		IsUrlValid func(toTest string) bool
		args1 int
		args2 int
		out1 *models.Url
		out2 error
	} {
		{
			name: "success no error",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args1: oldUrl.ID,
			args2: oldUrl.UserID,
			out1: oldUrl,
			out2: nil,
		},
		{
			name: "should error when url not found",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args1: oldUrl.ID,
			args2: oldUrl.UserID,
			out1: &models.Url{},
			out2: shared.ErrUrlNotFound,
		},
		{
			name: "should error when user not allowed to edit url",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return false
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args1: oldUrl.ID,
			args2: oldUrl.UserID,
			out1: &models.Url{},
			out2: shared.ErrForbiddenToAccess,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.IUrlRepository)

			// mock helper function
			services.IsUserAllowedToEditUrl = tc.IsUserAllowedToEditUrl
			services.IsUrlValid = tc.IsUrlValid

			if(tc.name == "should error when url not found") {
				mockRepo.On("GetUrlById", oldUrl.ID).Return(&models.Url{ID: 0}, nil)
			} else {
				mockRepo.On("GetUrlById", oldUrl.ID).Return(oldUrl, nil)
			}

			urlService := services.NewUrlService(mockRepo)
			got, gotErr := urlService.GetUrlUserById(tc.args1, tc.args2)
			assert.Equal(t, tc.out1, got)
			assert.Equal(t, tc.out2, gotErr)
		})
	}
}

func TestCreateUrl(t *testing.T) {
	services.Now = func() time.Time {
		return time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	generatedLink := "generatedLink"
	urlRequest := &models.Url{
		Title: "Test Title",
		ShortLink: generatedLink,
		OriginalLink: "http://localhost:8080/test",
		HitCounter:   0,
		UserID:       1,
		CreatedAt:    services.Now(),
		UpdatedAt:    services.Now(),
	}
	urlResponse := &models.Url{
		ID: 1,
		Title: "Test Title",
		ShortLink: generatedLink,
		OriginalLink: "http://localhost:8080/test",
		HitCounter:   0,
		UserID:       1,
		CreatedAt:    services.Now(),
		UpdatedAt:    services.Now(),
	}

	cases := []struct{
		name string
		IsUserAllowedToEditUrl func(userID int, userIDInUrl int) bool
		IsUrlValid func(toTest string) bool
		GenerateLink func() (string, error)
		args *dto.CreateUrlRequest
		out1 *dto.CreateUrlResponse
		out2 error
	} {
		{
			name: "success no error",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			GenerateLink: func() (string, error) {
				return generatedLink, nil
			},
			args: &dto.CreateUrlRequest{
				Title: urlRequest.Title,
				OriginalLink: urlRequest.OriginalLink,
				UserID: urlRequest.UserID,
			},
			out1: &dto.CreateUrlResponse{
				ID: urlResponse.ID,
				Title: urlResponse.Title,
				ShortLink: urlResponse.ShortLink,
				OriginalLink: urlResponse.OriginalLink,
				HitCounter: urlResponse.HitCounter,
				UserID: urlResponse.UserID,
				CreatedAt: urlResponse.CreatedAt,
				UpdatedAt: urlResponse.UpdatedAt,
			},
			out2: nil,
		},
		{
			name: "should error when original link not valid",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return false
			},
			GenerateLink: func() (string, error) {
				return generatedLink, nil
			},
			args: &dto.CreateUrlRequest{
				Title: urlRequest.Title,
				OriginalLink: "google",
				UserID: urlRequest.UserID,
			},
			out1: &dto.CreateUrlResponse{},
			out2: shared.ErrOriginalUrlNotValid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.IUrlRepository)
			mockRepo.On("CreateUrl", urlRequest).Return(urlResponse, nil)
			urlService := services.NewUrlService(mockRepo)

			// mock helper function
			services.IsUrlValid = tc.IsUrlValid
			services.GenerateLink = tc.GenerateLink

			got, gotErr := urlService.CreateUrl(tc.args)
			assert.Equal(t, tc.out1, got)
			assert.Equal(t, tc.out2, gotErr)
		})
	}
	
}

func TestUpdateUrl(t *testing.T) {
	services.Now = func() time.Time {
		return time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	generatedLink := "generatedLink"
	oldUrl := &models.Url{
		ID: 1,
		Title: "Test Title",
		ShortLink: generatedLink,
		OriginalLink: "http://localhost:8080/test",
		HitCounter:   0,
		UserID:       1,
		CreatedAt:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    services.Now(),
	}
	newUrl := &models.Url{
		ID: 1,
		Title: "Test Title",
		ShortLink: "newShortLink",
		OriginalLink: "http://localhost:8080/test",
		HitCounter:   0,
		UserID:       1,
		CreatedAt:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    services.Now(),
	}

	cases := []struct{
		name string
		IsUserAllowedToEditUrl func(userID int, userIDInUrl int) bool
		IsUrlValid func(toTest string) bool
		args *dto.UpdateUrlRequest
		out1 *dto.UpdateUrlResponse
		out2 error
	} {
		{
			name: "success no error",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args: &dto.UpdateUrlRequest{
				Title: newUrl.Title,
				ShortLink: newUrl.ShortLink,
				OriginalLink: newUrl.OriginalLink,
				UserID: newUrl.UserID,
			},
			out1: &dto.UpdateUrlResponse{
				ID: newUrl.ID,
				Title: newUrl.Title,
				ShortLink: newUrl.ShortLink,
				OriginalLink: newUrl.OriginalLink,
				HitCounter: newUrl.HitCounter,
				UserID: newUrl.UserID,
				CreatedAt: newUrl.CreatedAt,
				UpdatedAt: newUrl.UpdatedAt,
			},
			out2: nil,
		},
		{
			name: "should error when url not found",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args: &dto.UpdateUrlRequest{
				Title: newUrl.Title,
				ShortLink: newUrl.ShortLink,
				OriginalLink: newUrl.OriginalLink,
				UserID: newUrl.UserID,
			},
			out1: &dto.UpdateUrlResponse{},
			out2: shared.ErrUrlNotFound,
		},
		{
			name: "should error when original link not valid",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return false
			},
			args: &dto.UpdateUrlRequest{
				Title: newUrl.Title,
				ShortLink: newUrl.ShortLink,
				OriginalLink: newUrl.OriginalLink,
				UserID: newUrl.UserID,
			},
			out1: &dto.UpdateUrlResponse{},
			out2: shared.ErrOriginalUrlNotValid,
		},
		{
			name: "should error when new short link exist",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args: &dto.UpdateUrlRequest{
				Title: newUrl.Title,
				ShortLink: newUrl.ShortLink,
				OriginalLink: newUrl.OriginalLink,
				UserID: newUrl.UserID,
			},
			out1: &dto.UpdateUrlResponse{},
			out2: shared.ErrUrlShortLinkAlreadyExist,
		},
		{
			name: "should error when user not allowed to edit url",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return false
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args: &dto.UpdateUrlRequest{
				Title: newUrl.Title,
				ShortLink: newUrl.ShortLink,
				OriginalLink: newUrl.OriginalLink,
				UserID: newUrl.UserID,
			},
			out1: &dto.UpdateUrlResponse{},
			out2: shared.ErrForbiddenToAccess,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.IUrlRepository)
			mockRepo.On("UpdateUrl", oldUrl).Return(newUrl, nil)

			// mock helper function
			services.IsUserAllowedToEditUrl = tc.IsUserAllowedToEditUrl
			services.IsUrlValid = tc.IsUrlValid

			if(tc.name == "should error when new short link exist") {
				mockRepo.On("GetUrlByShortLink", tc.args.ShortLink).Return(&models.Url{
					ID: 4,
					Title: "Test Title",
					ShortLink: "newShortLink",
					OriginalLink: "http://localhost:8080/test",
					HitCounter:   0,
					UserID:       1,
					CreatedAt:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:    services.Now(),
				}, nil)
			} else {
				mockRepo.On("GetUrlByShortLink", tc.args.ShortLink).Return(&models.Url{}, nil)
			}

			if(tc.name == "should error when url not found") {
				mockRepo.On("GetUrlById", oldUrl.ID).Return(&models.Url{ID: 0}, nil)
			} else {
				mockRepo.On("GetUrlById", oldUrl.ID).Return(oldUrl, nil)
			}

			urlService := services.NewUrlService(mockRepo)
			got, gotErr := urlService.UpdateUrl(newUrl.ID, tc.args)
			assert.Equal(t, tc.out1, got)
			assert.Equal(t, tc.out2, gotErr)
		})
	}
}

func TestDeleteUrl(t *testing.T) {
	oldUrl := &models.Url{
		ID: 1,
		Title: "Test Title",
		ShortLink: "generatedLink",
		OriginalLink: "http://localhost:8080/test",
		HitCounter:   0,
		UserID:       1,
		CreatedAt:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	cases := []struct{
		name string
		IsUserAllowedToEditUrl func(userID int, userIDInUrl int) bool
		IsUrlValid func(toTest string) bool
		args1 int
		args2 int
		out1 *dto.UpdateUrlResponse
		out2 error
	} {
		{
			name: "success no error",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args1: oldUrl.ID,
			args2: oldUrl.UserID,
			out1: &dto.UpdateUrlResponse{
				ID: oldUrl.ID,
				Title: oldUrl.Title,
				ShortLink: oldUrl.ShortLink,
				OriginalLink: oldUrl.OriginalLink,
				HitCounter: oldUrl.HitCounter,
				UserID: oldUrl.UserID,
				CreatedAt: oldUrl.CreatedAt,
				UpdatedAt: oldUrl.UpdatedAt,
			},
			out2: nil,
		},
		{
			name: "should error when url not found",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return true
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args1: oldUrl.ID,
			args2: oldUrl.UserID,
			out1: &dto.UpdateUrlResponse{},
			out2: shared.ErrUrlNotFound,
		},
		{
			name: "should error when user not allowed to edit url",
			IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
				return false
			},
			IsUrlValid: func(toTest string) bool {
				return true
			},
			args1: oldUrl.ID,
			args2: oldUrl.UserID,
			out1: &dto.UpdateUrlResponse{},
			out2: shared.ErrForbiddenToAccess,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.IUrlRepository)
			mockRepo.On("DeleteUrl", oldUrl).Return(oldUrl, nil)

			// mock helper function
			services.IsUserAllowedToEditUrl = tc.IsUserAllowedToEditUrl
			services.IsUrlValid = tc.IsUrlValid

			if(tc.name == "should error when url not found") {
				mockRepo.On("GetUrlById", oldUrl.ID).Return(&models.Url{ID: 0}, nil)
			} else {
				mockRepo.On("GetUrlById", oldUrl.ID).Return(oldUrl, nil)
			}

			urlService := services.NewUrlService(mockRepo)
			got, gotErr := urlService.DeleteUrl(tc.args1, tc.args2)
			assert.Equal(t, tc.out1, got)
			assert.Equal(t, tc.out2, gotErr)
		})
	}
}