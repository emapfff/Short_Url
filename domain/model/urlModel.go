package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	ID          uint   `gorm:"unique;primaryKey;autoIncrement:true,"`
	OriginalUrl string `gorm:"unique"`
	ShortUrl    string `gorm:"size:10"`
}
