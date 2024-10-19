package repository

import (
	"gorm.io/gorm"
	"log"
	"test_ozon/domain/model"
)

func CreateUrlTable(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&model.Url{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SaveUrls(db *gorm.DB, originalUrl, shortUrl string) {
	db.Create(&model.Url{
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
	})
}

func GetOriginalUrl(db *gorm.DB, shortUrl string) (*string, error) {
	var url model.Url
	result := db.Where("short_url = ?", shortUrl, &url)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &url.OriginalUrl, nil
}

func GetAllUrls(db *gorm.DB) (*[]model.Url, error) {
	var urls []model.Url
	result := db.Find(&urls)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &urls, nil
}
