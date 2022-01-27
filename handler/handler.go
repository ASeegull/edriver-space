package handler

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/middleware"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

type Users interface {
	InitUsersRoutes(e *echo.Group, mw middleware.Middleware)
	SignIn() echo.HandlerFunc
	SignUp() echo.HandlerFunc
	SignOut() echo.HandlerFunc
	RefreshTokens() echo.HandlerFunc
	AddDriverLicence() echo.HandlerFunc
	AddVehicle() echo.HandlerFunc
	GetFines() echo.HandlerFunc
	PayAllFines() echo.HandlerFunc
	PayFine() echo.HandlerFunc
}

// Uploader provides methods to upload fines on the server
type Uploader interface {
	InitUploaderRoutes(e *echo.Group, mw middleware.Middleware)
	UploadXMLFines() echo.HandlerFunc
	UploadExcel() echo.HandlerFunc
}

type Drivers interface {
	InitDriversRoutes(e *echo.Group)
	CreateDriver() echo.HandlerFunc
	GetDriver() echo.HandlerFunc
	GetDrivers() echo.HandlerFunc
	DeleteDriver() echo.HandlerFunc
}

type Cars interface {
	InitCarsRoutes(e *echo.Group)
	CreateCar() echo.HandlerFunc
	GetCar() echo.HandlerFunc
	GetCars() echo.HandlerFunc
	DeleteCar() echo.HandlerFunc
}

type Police interface {
	InitPoliceRoutes(e *echo.Group, mw middleware.Middleware)
	GetFines() echo.HandlerFunc
	RemoveFine() echo.HandlerFunc
	GetFine() echo.HandlerFunc
}

// Handlers stores all handlers
type Handlers struct {
	Users   Users
	Upload  Uploader
	Drivers Drivers
	Cars    Cars
	Police  Police
}

// NewHandlers returns a pointer to new Handlers
func NewHandlers(services *service.Services, cfg *config.Config) *Handlers {
	return &Handlers{
		Users:   NewUsersHandlers(services.Users, cfg),
		Upload:  NewUploadHandler(services.Uploader, cfg),
		Drivers: NewDriverHandlers(services.Drivers, cfg),
		Cars:    NewCarsHandlers(services.Cars, cfg),
		Police:  NewPoliceHandler(services.Police, cfg),
	}
}

func (h *Handlers) InitRoutes(e *echo.Group, mw middleware.Middleware) {
	h.Users.InitUsersRoutes(e, mw)
	h.Upload.InitUploaderRoutes(e, mw)
	h.Drivers.InitDriversRoutes(e)
	h.Cars.InitCarsRoutes(e)
	h.Police.InitPoliceRoutes(e, mw)
}
