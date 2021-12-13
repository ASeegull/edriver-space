package utils

import (
	"github.com/ASeegull/edriver-space/config"
	"net/http"
	"time"
)

func CreateCookie(refreshToken string, cfg *config.Config) *http.Cookie {
	return &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: cfg.Cookie.HTTPOnly,
		SameSite: 0,
		Expires:  time.Now().Add(time.Duration(cfg.Cookie.Expire) * time.Minute),
	}
}

func DeleteCookie(name string) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
}
