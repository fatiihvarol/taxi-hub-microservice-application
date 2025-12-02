package dtos

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role,omitempty"` // opsiyonel: admin, driver, customer
}
