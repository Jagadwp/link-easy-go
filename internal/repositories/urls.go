package repositories

import (
	"errors"
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
	"github.com/Jagadwp/link-easy-go/internal/shared"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (u *UrlRepository) GetUrlsByUsername(username string) ([]models.Url, error) {
	var urls []models.Url
	u.db.Where("user_id = ?", username).Find(&urls)
	return urls, nil
}

func (u *UrlRepository) CreateShortUrl(title, originalLink, shortLink string, userID *int) (*models.Url, error) {
	//UserID itu pointer, tapi yg masuk value aslinya (int) 👍🏻
	var url models.Url = models.Url{
		Title:        title,
		ShortLink:    shortLink,
		OriginalLink: originalLink,
		HitCounter:   0,
		UserID:       userID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := u.db.Create(&url).Error; err != nil {
		var perr *pgconn.PgError
		errors.As(err, &perr)
		if perr.Code == shared.CODE_ERROR_DUPLICATE_KEY {
			return &models.Url{}, shared.ErrUrlShortLinkAlreadyExist
		}
		return &models.Url{}, err
	}

	return &url, nil
}

func (u *UrlRepository) InsertUrl(title string, shortLink string, originalLink string, userID *int) (*models.Url, error) {
	var url models.Url = models.Url{
		ShortLink:    shortLink,
		Title:        title,
		OriginalLink: originalLink,
		HitCounter:   0,
		UserID:       userID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := u.db.Create(&url).Error; err != nil {
		var perr *pgconn.PgError
		errors.As(err, &perr)
		if perr.Code == shared.CODE_ERROR_DUPLICATE_KEY {
			return &models.Url{}, shared.ErrUrlShortLinkAlreadyExist
		}
		return &models.Url{}, err
	}

	return &url, nil
}

func (u *UrlRepository) GetAllUrlsByUserID(userID int) (*[]models.Url, error) {
	var urls []models.Url

	if err := u.db.Where("user_id = ?", userID).Find(&urls).Error; err != nil {
		return &[]models.Url{}, err
	}

	return &urls, nil
}

func (u *UrlRepository) GetUrlById(id int) (*models.Url, error) {
	var url models.Url

	if err := u.db.Where("id = ?", id).Find(&url).Error; err != nil {
		return &models.Url{}, err
	}

	return &url, nil
}

func (u *UrlRepository) GetUrlByShortLink(shortLink string) (*models.Url, error) {
	var url models.Url

	if err := u.db.Where("short_link = ?", shortLink).Find(&url).Error; err != nil {
		return &models.Url{}, err
	}

	u.IncrementHitCounter(&url)

	return &url, nil
}

func (u *UrlRepository) IncrementHitCounter(url *models.Url) (*models.Url, error) {
	query := u.db.Model(url).Update("hit_counter", url.HitCounter+1)

	if err := query.Error; err != nil {
		return &models.Url{}, err
	}

	return url, nil
}

func (u *UrlRepository) UpdateUrl(url *models.Url) (*models.Url, error) {
	if err := u.db.Save(url).Error; err != nil {
		return &models.Url{}, err
	}

	return url, nil
}

func (u *UrlRepository) DeleteUrl(url *models.Url) (*models.Url, error) {
	if err := u.db.Delete(url).Error; err != nil {
		return &models.Url{}, err
	}

	return url, nil
}
