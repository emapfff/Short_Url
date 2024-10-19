package dto

type SaveUrl struct {
	OriginalUrl string `json:"original_url"`
}

type GetUrl struct {
	ShortUrl string `json:"short_url"`
}
