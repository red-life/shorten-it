package http

type ErrorResponse struct {
	err error
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required,http_url"`
}

type ShortenResponse struct {
	Key string `json:"key"`
}

type RedirectRequest struct {
	Key string `uri:"key" binding:"required,ascii"`
}
