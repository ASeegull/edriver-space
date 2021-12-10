package repository

import (
	"context"
	"encoding/json"
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/ASeegull/edriver-space/internal/session"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

const basePrefix = "api-session:"
type sessionRepo struct {
	redisClient *redis.Client
}

func NewSessionRepo(redisClient *redis.Client) session.SessRepository {
	return &sessionRepo{redisClient: redisClient}
}

func (r sessionRepo) CreateSession(ctx context.Context, sess *models.Session, expire int) (string, error) {
	sess.SessionID = uuid.New().String()

	sessBytes, err := json.Marshal(&sess)
	if err != nil {
		return "", err
	}

	if err := r.redisClient.Set(ctx, sess.SessionID, sessBytes, time.Second*time.Duration(expire)).Err(); err != nil {
		return "", err
	}

	return sess.SessionID, nil
}

func (r sessionRepo) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {

	sessBytes, err := r.redisClient.Get(ctx, sessionID).Bytes()
	if err != nil {
		return nil, err
	}

	sess := &models.Session{}
	if err := json.Unmarshal(sessBytes, &sess); err != nil {
		return nil, err
	}
	return sess, nil
}

func (r sessionRepo) DeleteByID(ctx context.Context, sessionID string) error {
	if err := r.redisClient.Del(ctx, sessionID).Err(); err != nil {
		return err
	}
	return nil
}
