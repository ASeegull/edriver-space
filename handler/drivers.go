package handler

import (
	"net/http"

	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

type DriversHandlers struct {
	driverService service.Drivers
	cfg           *config.Config
}

func NewDriverHandlers(driverService service.Drivers, cfg *config.Config) *DriversHandlers {
	return &DriversHandlers{
		driverService: driverService,
		cfg:           cfg,
	}
}

func (driversHandler *DriversHandlers) InitDriversRoutes(e *echo.Group) {
	driversRouters := e.Group("/drivers")

	driversRouters.POST("/", driversHandler.CreateDriver())
	driversRouters.GET("/:id", driversHandler.GetDriver())
	driversRouters.GET("/", driversHandler.GetDrivers())
	driversRouters.DELETE("/:id", driversHandler.DeleteDriver())
}

func (driversHandler *DriversHandlers) CreateDriver() echo.HandlerFunc {
	return func(c echo.Context) error {
		var driver model.Driver
		if err := c.Bind(&driver); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdDriver, err := driversHandler.driverService.CreateDriver(c.Request().Context(), &driver)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, createdDriver)
	}
}

func (driversHandler *DriversHandlers) GetDriver() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, ok := c.Get("id").(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "driver id is not present in context")
		}

		driver, err := driversHandler.driverService.GetDriver(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, driver)
	}
}

func (driversHandler *DriversHandlers) GetDrivers() echo.HandlerFunc {
	return func(c echo.Context) error {
		drivers, err := driversHandler.driverService.GetDrivers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, drivers)
	}
}

func (driversHandler *DriversHandlers) DeleteDriver() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, ok := c.Get("id").(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "driver id is not present in context")
		}

		err := driversHandler.driverService.DeleteDriver(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
