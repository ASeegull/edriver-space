package server

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/config"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo        *echo.Echo
	db          *sql.DB
	redisClient *redis.Client
	cfg         *config.Config
}

func NewServer(db *sql.DB, redisClient *redis.Client, cfg *config.Config) *Server {
	return &Server{
		echo:        echo.New(),
		db:          db,
		redisClient: redisClient,
		cfg:         cfg,
	}
}

func (s *Server) Run() error {

	port := s.cfg.Server.Port

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	if err := s.echo.Start(port); err != nil {
		return err
	}
	return s.echo.Shutdown(context.Background())
}
