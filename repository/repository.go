package repository

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/model"
	"github.com/go-redis/redis/v8"
	"time"
)

type Auth interface {
	GetUserByCredentials(ctx context.Context, login, password string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
}

type Sessions interface {
	SetSession(ctx context.Context, refreshToken, userId string, ttl time.Duration) error
	GetSessionById(ctx context.Context, sessionId string) (*string, error)
	DeleteSession(ctx context.Context, sessionId string) error
}

type Repositories struct {
	Auth     Auth
	Sessions Sessions
}

func NewRepositories(postgres *sql.DB, redis *redis.Client) *Repositories {
	return &Repositories{
		Auth:     NewAuthRepos(postgres),
		Sessions: NewSessionsRepos(redis),
	}
}
