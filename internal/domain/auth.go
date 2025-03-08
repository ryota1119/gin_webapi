package domain

import "time"

// AuthToken ユーザーに発行された認証トークン
type AuthToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
