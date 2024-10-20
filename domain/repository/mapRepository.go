package repository

import (
	"errors"
	"log"
)

type MapLinkRepo struct {
	store map[string]string
}

func CreateUrlMap() {
	log.Println("Map is in use")
	Repo = &MapLinkRepo{store: make(map[string]string)}
}

func (repo *MapLinkRepo) SaveUrls(originalUrl, shortUrl string) {
	repo.store[originalUrl] = shortUrl
}

func (repo *MapLinkRepo) GetOriginalUrl(shortUrl string) (*string, error) {
	for key, value := range repo.store {
		if value == shortUrl {
			return &key, nil
		}
	}
	return nil, errors.New("does not exist such url")
}

func (repo *MapLinkRepo) CheckExistOriginalUrl(originalUrl string) bool {
	_, exists := repo.store[originalUrl]
	return exists
}
