package service

import (
	"context"

	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/repository"
)

type carService struct {
	cfg     *config.Config
	carRepo repository.Cars
}

func NewCarService(cfg *config.Config, carRepo repository.Cars) *carService {
	return &carService{cfg: cfg, carRepo: carRepo}
}

func (c *carService) CreateCar(ctx context.Context, car *model.Car) (*model.Car, error) {
	createdUser, err := c.carRepo.CreateCar(ctx, car)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (c *carService) GetCar(ctx context.Context, id string) (*model.Car, error) {
	getCar, err := c.carRepo.GetCar(ctx, id)
	if err != nil {
		return nil, err
	}

	return getCar, nil
}

func (c *carService) GetCars(ctx context.Context) (*[]model.Car, error) {
	cars, err := c.carRepo.GetCars(ctx)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *carService) DeleteCar(ctx context.Context, id string) error {
	err := c.carRepo.DeleteCar(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
