package repository

import (
	"gorm.io/gorm"
	"log"
	"test_ozon/domain"
	"test_ozon/domain/model"
)

var Repo domain.UrlModel

type DbLinkRepo struct {
	db *gorm.DB
}

func CreateUrlTable(db *gorm.DB) {
	err := db.AutoMigrate(&model.Url{})
	if err != nil {
		log.Fatal(err)
	}
	Repo = &DbLinkRepo{
		db: db,
	}
}

func (repo *DbLinkRepo) SaveUrls(originalUrl, shortUrl string) {
	repo.db.Create(&model.Url{
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
	})
}

func (repo *DbLinkRepo) GetOriginalUrl(shortUrl string) (*string, error) {
	var url model.Url
	result := repo.db.Where("short_url = ?", shortUrl).First(&url)

	if result.Error != nil {
		return nil, result.Error
	}
	return &url.OriginalUrl, nil
}

func (repo *DbLinkRepo) CheckExistOriginalUrl(originalUrl string) bool {
	var count int64
	err := repo.db.Model(model.Url{}).Where("original_url = ?", originalUrl).Limit(1).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	return count > 0
}
