package dtos

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResponse struct {
	Errors []ValidationError `json:"errors"`
}
