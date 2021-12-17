package repository

import (
	"context"
	"github.com/ASeegull/edriver-space/models"
	"github.com/go-redis/redis/v8"
	"time"
)

type SessionsRepos struct {
	client *redis.Client
}

func NewSessionsRepos(client *redis.Client) *SessionsRepos {
	return &SessionsRepos{
		client: client,
	}
}

func (s *SessionsRepos) SetSession(ctx context.Context, refreshToken, userId string, ttl time.Duration) error {
	//sessBytes, err := json.Marshal(userId)
	//if err != nil {
	//	return err
	//}

	return s.client.Set(ctx, refreshToken, userId, ttl).Err()
}

func (s *SessionsRepos) GetSessionById(ctx context.Context, sessionId string) (*string, error) {

	userId, err := s.client.Get(ctx, sessionId).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, models.ErrSessionNotFound
		}

		return nil, err
	}
	return &userId, nil
}

func (s *SessionsRepos) DeleteSession(ctx context.Context, sessionId string) error {
	return s.client.Del(ctx, sessionId).Err()
}
