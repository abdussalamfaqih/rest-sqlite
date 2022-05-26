package domain

// LoginRequest ...
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type TokenData struct {
	TokenType string `json:"type"`
	Token     string `json:"token"`
	ExpiresAt int    `json:"expires_at"`
}
