package middleware

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/session"
)

type Middleware struct {
	sessUC session.UCSession
	authUC auth.UseCase
	cfg    *config.Config
}

func NewMiddleware(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config) *Middleware {
	return &Middleware{
		sessUC: sessUC,
		authUC: authUC,
		cfg:    cfg,
	}
}
