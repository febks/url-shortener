package models

type URLRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    DataPayload `json:"data"`
}

type DataPayload struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ResultURL   string `json:"result_url"`
}
