package main

import (
	"fmt"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	_ "github.com/ASeegull/edriver-space/migrations"
	"github.com/ASeegull/edriver-space/server"
	"github.com/ASeegull/edriver-space/storage"
	"github.com/pressly/goose/v3"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		logger.LogFatal(err)
	}

	fmt.Println(cfg)

	postgres, err := storage.NewPostgresDB(cfg)
	if err != nil {
		logger.LogFatal(err)
	}

	defer func() {
		if conErr := postgres.Close(); conErr != nil {
			logger.LogFatal(err)
		}
	}()

	if err := goose.Up(postgres, "."); err != nil {
		logger.LogFatal(err)
	}

	redis := storage.NewRedisClient(cfg)

	s := server.NewServer(postgres, redis, cfg)
	if err := s.Run(); err != nil {
		logger.LogFatal(err)
	}
}
