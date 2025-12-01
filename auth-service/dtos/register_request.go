package dtos

type RegisterRequest struct {
    Email    string `json:"email" example:"fatih@example.com"`
    Password string `json:"password" example:"123456"`
}
