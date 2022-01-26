package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ASeegull/edriver-space/model"
	"github.com/go-redis/redis/v8"
)

type Users interface {
	GetUserByCredentials(ctx context.Context, email, password string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	CreateUser(ctx context.Context, newUser model.User) (string, error)
	CreateUserPolice(ctx context.Context, newUser model.User) (string, error)
	GetDriverLicence(ctx context.Context, individualTaxNumber string) (string, error)
	UpdateUserDriverLicence(ctx context.Context, userId, licenceNumber string) error
	GetCarsFines(ctx context.Context, userId string) ([]model.CarsFine, error)
	GetDriversFines(ctx context.Context, userId string) ([]model.DriversFine, error)
}

type Sessions interface {
	SetSession(ctx context.Context, refreshToken, userId string, ttl time.Duration) error
	GetSessionById(ctx context.Context, sessionId string) (*string, error)
	DeleteSession(ctx context.Context, sessionId string) error
}

type Cars interface {
	GetCar(ctx context.Context, id string) (*model.Car, error)
	GetCars(ctx context.Context) (*[]model.Car, error)
	CreateCar(ctx context.Context, car *model.Car) (*model.Car, error)
	DeleteCar(ctx context.Context, id string) error
}

type Drivers interface {
	GetDriver(ctx context.Context, id string) (*model.Driver, error)
	GetDrivers(ctx context.Context) (*[]model.Driver, error)
	CreateDriver(ctx context.Context, driver *model.Driver) (*model.Driver, error)
	DeleteDriver(ctx context.Context, id string) error
}

type DriverFines interface {
	GetDriverFines(ctx context.Context, licence string) ([]model.DriversFine, error)
	GetDriverFine(ctx context.Context, licence string) (*model.DriversFine, error)
	AddDriverFine(ctx context.Context, fine *model.DriversFine) error
	DeleteDriverFine(ctx context.Context, fineNum string) error
}

type CarFines interface {
	GetCarFines(ctx context.Context, regNum string) ([]model.CarsFine, error)
	GetCarFine(ctx context.Context, regNum string) (*model.CarsFine, error)
	AddCarFine(ctx context.Context, fine *model.CarsFine) error
	DeleteCarFine(ctx context.Context, fineNum string) error
}

type Repositories struct {
	Users       Users
	Sessions    Sessions
	Cars        Cars
	Drivers     Drivers
	DriverFines DriverFines
	CarFines    CarFines
}

func NewRepositories(postgres *sql.DB, redis *redis.Client) *Repositories {
	return &Repositories{
		Users:       NewUsersRepos(postgres),
		Sessions:    NewSessionsRepos(redis),
		Cars:        NewCarsRepos(postgres),
		Drivers:     NewDriversRepos(postgres),
		DriverFines: NewDriverFinesRep(postgres),
		CarFines:    NewCarFinesRep(postgres),
	}
}
