package repositories

import (
	"github.com/Jagadwp/link-easy-go/internal/models"
	"gorm.io/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (u *UrlRepository) GetUrlsByUsername(username string) (*[]models.Url, error) {
	var urls []models.Url
	u.db.Where("user_id = ?", username).Find(&urls)
	return &urls, nil
}

func (u *UrlRepository) GetUrlsByUserID(userID int) (*[]models.Url, error) {
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

	//TODO harusnya ini gaboleh disini
	// u.IncrementHitCounter(&url)

	return &url, nil
}

func (u *UrlRepository) IncrementHitCounter(url *models.Url) (*models.Url, error) {
	query := u.db.Model(url).Update("hit_counter", url.HitCounter+1)

	if err := query.Error; err != nil {
		return &models.Url{}, err
	}

	return url, nil
}

func (u *UrlRepository) CreateUrl(url *models.Url) (*models.Url, error) {
	if err := u.db.Create(url).Error; err != nil {
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
