package usecase

import (
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/ASeegull/edriver-space/internal/session"
	jwt "github.com/ASeegull/edriver-space/pkg/auth"
	"time"
)

type sessionUC struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
	jwtManager  *jwt.JWTManager
}

func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config, jwtManager *jwt.JWTManager) session.UCSession {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg, jwtManager: jwtManager}
}

func (u *sessionUC) CreateSession(ctx context.Context, userId string, refreshToken string, ttl time.Duration) error {

	return u.sessionRepo.CreateSession(ctx, userId, refreshToken, ttl)
}

func (u *sessionUC) DeleteByID(ctx context.Context, sessionID string) error {

	return u.sessionRepo.DeleteByID(ctx, sessionID)
}

func (u *sessionUC) GetSessionByID(ctx context.Context, sessionID string) (*models.RefreshSession, error) {

	return u.sessionRepo.GetSessionByID(ctx, sessionID)
}

func (u *sessionUC) RefreshSession(ctx context.Context, sessionID string) (*models.JWTTokens, error) {

	refreshSession, err := u.GetSessionByID(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	if err := u.DeleteByID(ctx, sessionID); err != nil {
		return nil, nil
	}

	newRefreshJWTToken, err := u.jwtManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	refreshTTL := time.Duration(u.cfg.Tokens.RefreshTokenTTL) * time.Minute
	if err := u.CreateSession(ctx, refreshSession.UserId, newRefreshJWTToken, refreshTTL); err != nil {
		return nil, err
	}

	accessTTL := time.Duration(u.cfg.Tokens.AccessTokenTTL) * time.Minute
	newAccessJWTToken, err := u.jwtManager.NewJWT(refreshSession.UserId, accessTTL)
	if err != nil {
		return nil, err
	}

	return &models.JWTTokens{
		AccessToken:  newAccessJWTToken,
		RefreshToken: newRefreshJWTToken,
	}, nil
}
