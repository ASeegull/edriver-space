package handler

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

type Auth interface {
	SignIn() echo.HandlerFunc
	SignUp() echo.HandlerFunc
	SignOut() echo.HandlerFunc
	RefreshTokens() echo.HandlerFunc
}

type Handlers struct {
	Auth Auth
}

func NewHandlers(services *service.Services, cfg *config.Config) *Handlers {
	return &Handlers{
		Auth: NewAuthHandlers(services.Auth, cfg),
	}
}
