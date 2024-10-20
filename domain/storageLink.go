package domain

type UrlModel interface {
	SaveUrls(originalUrl, shortUrl string)
	GetOriginalUrl(shortUrl string) (*string, error)
	CheckExistOriginalUrl(originalUrl string) bool
}
