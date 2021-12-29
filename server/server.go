package server

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/config"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo     *echo.Echo
	postgres *sql.DB
	redis    *redis.Client
	cfg      *config.Config
}

func NewServer(postgres *sql.DB, redis *redis.Client, cfg *config.Config) *Server {
	return &Server{
		echo:     echo.New(),
		postgres: postgres,
		redis:    redis,
		cfg:      cfg,
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
