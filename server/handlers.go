package server

import (
	"github.com/ASeegull/edriver-space/handler"
	"github.com/ASeegull/edriver-space/middleware"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/pkg/hash"
	"github.com/ASeegull/edriver-space/pkg/validator"
	"github.com/ASeegull/edriver-space/repository"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	tokenManager, err := auth.NewManager(s.cfg.Server.JWTSecretKey)
	if err != nil {
		return err
	}
	hasher, err := hash.NewSHA256Hasher(s.cfg.Server.HashSalt)
	if err != nil {
		return err
	}

	repositories := repository.NewRepositories(s.postgres, s.redis)
	services := service.NewServices(repositories, tokenManager, hasher, s.cfg)
	handlers := handler.NewHandlers(services, s.cfg)
	middlewares := middleware.NewMiddlewares(tokenManager)

	// init validator
	e.Validator = validator.NewValidationUtil()

	// All routes
	v1 := e.Group("/api/v1")

	// init all routes
	handlers.InitRoutes(v1, middlewares)

	return nil
}
