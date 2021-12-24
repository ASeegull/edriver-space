package storage

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisHost,
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.DB,       // use default DB
	})

	return client
}
