package server

import (
	authHttp "github.com/ASeegull/edriver-space/internal/auth/delivery/http"
	authRepository "github.com/ASeegull/edriver-space/internal/auth/repository"
	authUseCase "github.com/ASeegull/edriver-space/internal/auth/usecase"
	"github.com/ASeegull/edriver-space/internal/middleware"
	sessRepository "github.com/ASeegull/edriver-space/internal/session/repository"
	sessUseCase "github.com/ASeegull/edriver-space/internal/session/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	authRepo := authRepository.NewAuthRepository(s.db)
	sessRepo := sessRepository.NewSessionRepo(s.redisClient)

	authUC := authUseCase.NewAuthUseCase(authRepo, s.cfg)
	sessUC := sessUseCase.NewSessionUseCase(sessRepo, s.cfg)

	authHandlers := authHttp.NewAuthHandlers(authUC, sessUC, s.cfg)

	mw := middleware.NewMiddleware(sessUC, authUC, s.cfg)

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
