package middleware

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/session"
	jwt "github.com/ASeegull/edriver-space/pkg/auth"
)

type Middleware struct {
	sessUC     session.UCSession
	authUC     auth.UseCase
	cfg        *config.Config
	jwtManager *jwt.JWTManager
}

func NewMiddleware(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, jwtManager *jwt.JWTManager) *Middleware {
	return &Middleware{
		sessUC: sessUC,
		authUC: authUC,
		cfg:    cfg,
		jwtManager: jwtManager,
	}
}
