package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/ASeegull/edriver-space/internal/session"
	"github.com/go-redis/redis/v8"
	"time"
)

const basePrefix = "refresh-token:"

type sessionRepo struct {
	redisClient *redis.Client
}

func NewSessionRepo(redisClient *redis.Client) session.SessRepository {
	return &sessionRepo{redisClient: redisClient}
}

func (r sessionRepo) CreateSession(ctx context.Context, userId string, refreshToken string, ttl time.Duration) error {
	//rsBytes, err := json.Marshal(&rs)
	//if err != nil {
	//	return err
	//}

	if err := r.redisClient.Set(ctx, KeyWithPrefix(refreshToken), []byte(userId), ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (r sessionRepo) GetSessionByID(ctx context.Context, sessionId string) (*models.RefreshSession, error) {

	sessBytes, err := r.redisClient.Get(ctx, KeyWithPrefix(sessionId)).Bytes()
	if err != nil {
		return nil, err
	}

	sess := &models.RefreshSession{}
	if err := json.Unmarshal(sessBytes, &sess); err != nil {
		return nil, err
	}
	return sess, nil
}

func (r sessionRepo) DeleteByID(ctx context.Context, sessionId string) error {
	if err := r.redisClient.Del(ctx, KeyWithPrefix(sessionId)).Err(); err != nil {
		return err
	}
	return nil
}

func KeyWithPrefix(id string) string {
	return fmt.Sprintf("%s %s", basePrefix, id)
}
