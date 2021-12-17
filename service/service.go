package service

import (
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/models"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/repository"
)

type Auth interface {
	SignIn(ctx context.Context, user UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, sessionId string) (Tokens, error)
	GetUserById(ctx context.Context, userId string) (*models.User, error)
	DeleteSession(ctx context.Context, sessionId string) error
}

type Services struct {
	Auth         Auth
	TokenManager auth.TokenManager
}

func NewServices(repos *repository.Repositories, tokenManager auth.TokenManager, cfg *config.Config) *Services {
	return &Services{
		Auth: NewAuthService(repos, tokenManager, cfg),
	}
}
