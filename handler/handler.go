package handler

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

// Auth provides authentication logic
type Auth interface {
	SignIn() echo.HandlerFunc
	SignUp() echo.HandlerFunc
	SignOut() echo.HandlerFunc
	RefreshTokens() echo.HandlerFunc
}

// Uploader provides methods to upload fines on the server
type Uploader interface {
	UploadXMLFines() echo.HandlerFunc
	UploadExcel() echo.HandlerFunc
}

// Handlers stores all handlers
type Handlers struct {
	Auth   Auth
	Upload Uploader
}

// NewHandlers returns a pointer to new Handlers
func NewHandlers(services *service.Services, cfg *config.Config) *Handlers {
	return &Handlers{
		Auth:   NewAuthHandlers(services.Auth, cfg),
		Upload: NewUploadHandler(services.Uploader, cfg),
	}
}
