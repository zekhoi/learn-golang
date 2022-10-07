package request

type CreateShortenRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
	CustomUrl   string `json:"custom_url"`
}

type GetShortenRequest struct {
	ShortUrl string `json:"short_url"  binding:"required"`
}
