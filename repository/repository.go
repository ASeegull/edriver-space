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

type ParkingFines interface {
	GetParkingFine(ctx context.Context, id string) (*model.ParkingFine, error)
	GetParkingFines(ctx context.Context) ([]model.ParkingFine, error)
	AddParkingFine(ctx context.Context, fine model.ParkingFine) error
	DeleteParkingFine(ctx context.Context, id string) error
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

type Repositories struct {
	Users        Users
	Sessions     Sessions
	Cars         Cars
	Drivers      Drivers
	ParkingFines ParkingFines
}

func NewRepositories(postgres *sql.DB, redis *redis.Client) *Repositories {
	return &Repositories{

		Users:        NewUsersRepos(postgres),
		Sessions:     NewSessionsRepos(redis),
		ParkingFines: NewParkingFinesRepos(postgres),
		Cars:         NewCarsRepos(postgres),
		Drivers:      NewDriversRepos(postgres),
	}
}
