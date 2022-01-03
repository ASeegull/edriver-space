package server

import (
	"github.com/ASeegull/edriver-space/handler"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/pkg/hash"
	"github.com/ASeegull/edriver-space/pkg/middleware"
	"github.com/ASeegull/edriver-space/repository"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	tokenManager, err := auth.NewManager("secret_key")
	if err != nil {
		return err
	}
	hasher, err := hash.NewSHA256Hasher("salt")
	if err != nil {
		return err
	}

	repositories := repository.NewRepositories(s.postgres, s.redis)
	services := service.NewServices(repositories, tokenManager, hasher, s.cfg)
	handlers := handler.NewHandlers(services, s.cfg)
	middlewares := middleware.NewCustomMiddlewares()

	// All routes
	v1 := e.Group("/api/v1")

	// Secure group
	secure := v1.Group("/secure", middlewares.Authorize.JWTAuthorization("police"))

	authGroup := v1.Group("/auth")
	// auth routes
	authGroup.POST("/sign-in", handlers.Auth.SignIn())
	authGroup.POST("/sign-out", handlers.Auth.SignOut())
	authGroup.POST("/sign-up", handlers.Auth.SignUp())

	authGroup.GET("/refresh-tokens", handlers.Auth.RefreshTokens())

	// Upload routes (secure access)
	uploadGroup := secure.Group("/upload")

	uploadGroup.POST("/XML", handlers.Upload.UploadXMLFines()) // Upload XML fines data to the server
	uploadGroup.POST("/Excel", handlers.Upload.UploadExcel())  // Upload Excel file with fines to the server

	return nil
}
