package httputils

type ApiResponseDto struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ApiErrorResponseDto struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}
