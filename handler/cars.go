package handler

import (
	"net/http"

	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

type CarsHandlers struct {
	carService service.Cars
	cfg        *config.Config
}

func NewCarsHandlers(carService service.Cars, cfg *config.Config) *CarsHandlers {
	return &CarsHandlers{
		carService: carService,
		cfg:        cfg,
	}
}

func (carsHandler *CarsHandlers) CreateCar() echo.HandlerFunc {
	return func(c echo.Context) error {
		var car model.Car
		if err := c.Bind(&car); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdCar, err := carsHandler.carService.CreateCar(c.Request().Context(), &car)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, createdCar)
	}
}

func (carsHandler *CarsHandlers) GetCar() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, ok := c.Get("id").(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "car id is not present in context")
		}

		car, err := carsHandler.carService.GetCar(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, car)
	}
}

func (carsHandler *CarsHandlers) GetCars() echo.HandlerFunc {
	return func(c echo.Context) error {
		cars, err := carsHandler.carService.GetCars(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, cars)
	}
}

func (carsHandler *CarsHandlers) DeleteCar() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, ok := c.Get("id").(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "car id is not present in context")
		}

		err := carsHandler.carService.DeleteCar(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
