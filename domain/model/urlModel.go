package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	OriginalUrl string `gorm:"unique"`
	ShortUrl    string `gorm:"size:10"`
}
