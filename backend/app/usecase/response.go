package usecase

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message    string      `json:"message"`
	Extensions interface{} `json:"extensions"`
}
