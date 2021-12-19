package repository

import (
	"context"
	"fmt"
	"github.com/ASeegull/edriver-space/models"
	"github.com/go-redis/redis/v8"
	"time"
)

const basePrefix = "session:"

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

	return s.client.Set(ctx, s.keyWithPrefix(refreshToken), userId, ttl).Err()
}

func (s *SessionsRepos) GetSessionById(ctx context.Context, sessionId string) (*string, error) {

	userId, err := s.client.Get(ctx, s.keyWithPrefix(sessionId)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, models.ErrSessionNotFound
		}

		return nil, err
	}
	return &userId, nil
}

func (s *SessionsRepos) DeleteSession(ctx context.Context, sessionId string) error {
	return s.client.Del(ctx, s.keyWithPrefix(sessionId)).Err()
}

func (s *SessionsRepos) keyWithPrefix(sessionId string) string {
	return fmt.Sprintf("%s %s", basePrefix, sessionId)
}
