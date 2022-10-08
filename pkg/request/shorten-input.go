package request

type CreateShortenRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
}

type GetShortenRequest struct {
	ShortUrl string `json:"short_url"  binding:"required"`
}
