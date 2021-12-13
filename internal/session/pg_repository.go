package session

import (
	"context"
	"github.com/ASeegull/edriver-space/internal/models"
	"time"
)

type SessRepository interface {
	CreateSession(ctx context.Context, userId string, refreshToken string, ttl time.Duration) error
	GetSessionByID(ctx context.Context, sessionID string) (*models.RefreshSession, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
