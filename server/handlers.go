package server

import (
	"github.com/ASeegull/edriver-space/handler"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/repository"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	tokenManager, err := auth.NewManager("secret_key")
	if err != nil {
		return err
	}

	repositories := repository.NewRepositories(s.postgres, s.redis)
	services := service.NewServices(repositories, tokenManager, s.cfg)
	handlers := handler.NewHandlers(services, s.cfg)

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	// auth routes
	authGroup.POST("/sign-in", handlers.Auth.SignIn())
	authGroup.POST("/sign-out", handlers.Auth.SignOut())
	authGroup.POST("/sign-up", handlers.Auth.SignUp())

	authGroup.GET("/refresh-tokens", handlers.Auth.RefreshTokens())

	return nil
}
