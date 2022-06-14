package repositories

import (
	"time"

	"github.com/Jagadwp/link-easy-go/internal/models"
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
	u.db.Where("created_by = ?", username).Find(&urls)
	return urls, nil
}

func (u *UrlRepository) InsertUrl(title string, shortLink string, originalLink string, userID int) (*models.Url, error) {
	var url models.Url = models.Url{
		ShortLink: shortLink,
		Title: title,
		OriginalLink: originalLink,
		HitCounter: 0,
		CreatedBy: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	if err:= u.db.Create(&url).Error; err != nil {
		return &models.Url{}, err
	}

	return &url, nil
}

func (u *UrlRepository) GetAllUrlsByUserID(userID int) (*[]models.Url, error) {
	var urls []models.Url
	
	if err:= u.db.Where("created_by = ?", userID).Find(&urls).Error; err != nil {
		return &[]models.Url{}, err
	}

	return &urls, nil
}

func (u *UrlRepository) GetUrlById(id int) (*models.Url, error) {
	var url models.Url
	
	if err:= u.db.Where("id = ?", id).Find(&url).Error; err != nil {
		return &models.Url{}, err
	}

	return &url, nil
}

func (u *UrlRepository) GetUrlByShortLink(shortLink string) (*models.Url, error) {
	var url models.Url
	
	if err:= u.db.Where("short_link = ?", shortLink).Find(&url).Error; err != nil {
		return &models.Url{}, err
	}

	return &url, nil
}

func (u *UrlRepository) UpdateUrl(url *models.Url) (*models.Url, error) {
	if err:= u.db.Save(url).Error; err != nil {
		return &models.Url{}, err
	}

	return url, nil
}

func (u *UrlRepository) DeleteUrl(url *models.Url) (*models.Url, error) {
	if err:= u.db.Delete(url).Error; err != nil {
		return &models.Url{}, err
	}

	return url, nil
}