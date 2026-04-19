package dto

type CreateUrlRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

type CreateUrlResponse struct {
	Id          int64  `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
}

type GetUrlByShortCodeResponse struct {
	Id          int64  `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
}
