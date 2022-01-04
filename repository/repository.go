package repository

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/model"
	"github.com/go-redis/redis/v8"
	"time"
)

type Users interface {
	GetUserByCredentials(ctx context.Context, email, password string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	CreateUser(ctx context.Context, newUser model.User) (string, error)
}

type Sessions interface {
	SetSession(ctx context.Context, refreshToken, userId string, ttl time.Duration) error
	GetSessionById(ctx context.Context, sessionId string) (*string, error)
	DeleteSession(ctx context.Context, sessionId string) error
}

type Uploader interface {
	GetFine(ctx context.Context, id string) (*model.ParkingFine, error)
	GetFines(ctx context.Context) ([]model.ParkingFine, error)
	AddFine(ctx context.Context, fine model.ParkingFine) error
	DeleteFine(ctx context.Context, id string) error
}

type Repositories struct {
	Users    Users
	Sessions Sessions
	Uploader Uploader
}

func NewRepositories(postgres *sql.DB, redis *redis.Client) *Repositories {
	return &Repositories{
		Users:    NewUsersRepos(postgres),
		Sessions: NewSessionsRepos(redis),
		Uploader: NewUploadRepos(postgres),
	}
}
