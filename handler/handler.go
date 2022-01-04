package handler

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

type Users interface {
	InitUsersRoutes(e *echo.Group)
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
	Users  Users
	Upload Uploader
}

// NewHandlers returns a pointer to new Handlers
func NewHandlers(services *service.Services, cfg *config.Config) *Handlers {
	return &Handlers{
		Users:  NewUsersHandlers(services.Users, cfg),
		Upload: NewUploadHandler(services.Uploader, cfg),
	}
}

func (h *Handlers) InitRoutes(e *echo.Group) {
	h.Users.InitUsersRoutes(e)
}
