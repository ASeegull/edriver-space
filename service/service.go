package service

import (
	"bytes"
	"context"

	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/pkg/hash"
	"github.com/ASeegull/edriver-space/repository"
)

type Users interface {
	SignUp(ctx context.Context, user UserSignUpInput) (Tokens, error)
	SignIn(ctx context.Context, user UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, sessionId string) (Tokens, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	DeleteSession(ctx context.Context, sessionId string) error
	AddDriverLicence(ctx context.Context, input AddDriverLicenceInput, userId string) error
	GetFines(ctx context.Context, userId string) (model.Fines, error)
}

type Uploader interface {
	XMLFinesService(ctx context.Context, data model.Data) error
	ReadFinesExcel(ctx context.Context, r *bytes.Reader) error
}

type Cars interface {
	CreateCar(ctx context.Context, car *model.Car) (*model.Car, error)
	GetCar(ctx context.Context, id string) (*model.Car, error)
	GetCars(ctx context.Context) (*[]model.Car, error)
	DeleteCar(ctx context.Context, id string) error
}

type Drivers interface {
	CreateDriver(ctx context.Context, driver *model.Driver) (*model.Driver, error)
	GetDriver(ctx context.Context, id string) (*model.Driver, error)
	GetDrivers(ctx context.Context) (*[]model.Driver, error)
	DeleteDriver(ctx context.Context, id string) error
}

type Services struct {
	Users    Users
	Uploader Uploader
	Cars     Cars
	Drivers  Drivers
}

func NewServices(repos *repository.Repositories, tokenManager auth.TokenManager, hasher hash.PasswordHasher, cfg *config.Config) *Services {
	return &Services{
		Users:    NewUsersService(repos.Users, repos.Sessions, tokenManager, hasher, cfg),
		Uploader: NewUploadService(repos, cfg),
		Cars:     NewCarService(cfg, repos.Cars),
		Drivers:  NewDriverService(cfg, repos.Drivers),
	}
}
