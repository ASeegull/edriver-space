package usecase

import (
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/ASeegull/edriver-space/internal/session"
)

type sessionUC struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
}

func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config) session.UCSession {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg}
}

func (u *sessionUC) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {

	return u.sessionRepo.CreateSession(ctx, session, expire)
}

func (u *sessionUC) DeleteByID(ctx context.Context, sessionID string) error {

	return u.sessionRepo.DeleteByID(ctx, sessionID)
}

func (u *sessionUC) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {

	return u.sessionRepo.GetSessionByID(ctx, sessionID)
}
