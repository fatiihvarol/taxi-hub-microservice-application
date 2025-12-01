package dtos

type LoginRequest struct {
    Email    string `json:"email" example:"fatih@example.com"`
    Password string `json:"password" example:"123456"`
}