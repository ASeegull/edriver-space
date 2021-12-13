package models

type Session struct {
	SessionID string `json:"session_id"`
	UserID    int    `json:"user_id"`
}

type RefreshSession struct {
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
}
