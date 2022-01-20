package handler

import (
	"errors"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/middleware"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type PoliceHandler struct {
	policeService service.Police
	cfg           *config.Config
}

// NewPoliceHandler returns pointer to new PoliceHandler
func NewPoliceHandler(policeService service.Police, cfg *config.Config) *PoliceHandler {
	return &PoliceHandler{policeService: policeService, cfg: cfg}
}

// InitPoliceRoutes initializes all police routes
func (ph *PoliceHandler) InitPoliceRoutes(e *echo.Group, mw middleware.Middleware) {
	// Police routes (police access only)
	police := e.Group("/police", mw.JWTAuthorization("police"))

	police.GET("/fines", ph.GetFines())     // Get all fines of given type with given user info
	police.GET("/fine", ph.GetFine())       // Get fine of given type with given fine number
	police.DELETE("/fine", ph.RemoveFine()) // Delete fine of given type with given fine number
}

// GetFines returns all car/driver type fines by users (registration number/licence)
func (ph *PoliceHandler) GetFines() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Get fine type from query parameter
		fineType := ctx.QueryParam("type")
		if fineType == "" {
			err := errors.New("fine type not specified")
			logger.LogErr(err)
			return ctx.JSON(http.StatusPreconditionRequired, err.Error())
		}
		// Get driver or car info from query parameter
		info := ctx.QueryParam("info")
		if info == "" {
			err := errors.New("user information is empty")
			logger.LogErr(err)
			return ctx.JSON(http.StatusPreconditionRequired, err.Error())
		}

		// Choose fines type to get
		switch strings.ToLower(fineType) {
		case "driver": // Get driver fines
			fines, err := ph.policeService.GetFinesDriverLicense(ctx.Request().Context(), info)
			if err != nil {
				logger.LogErr(err)
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}
			return ctx.JSON(http.StatusOK, fines) // response with the fines

		case "car": // Get car fines
			fines, err := ph.policeService.GetFinesCarRegNum(ctx.Request().Context(), info)
			if err != nil {
				logger.LogErr(err)
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}
			return ctx.JSON(http.StatusOK, fines) // response with the fines

		default:
			return ctx.JSON(http.StatusPreconditionFailed, "wrong fine type provided")
		}
	}
}

// GetFine returns driver/car type fine by fine number
func (ph *PoliceHandler) GetFine() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Get fine type from query parameter
		fineType := ctx.QueryParam("type")
		if fineType == "" {
			err := errors.New("fine type not specified")
			logger.LogErr(err)
			return ctx.JSON(http.StatusPreconditionRequired, err.Error())
		}
		// Get fine number from query parameter
		fineNum := ctx.QueryParam("num")
		if fineNum == "" {
			err := errors.New("fine number is empty")
			logger.LogErr(err)
			return ctx.JSON(http.StatusPreconditionRequired, err.Error())
		}

		// Choose fine type to get
		switch strings.ToLower(fineType) {
		case "driver": // Get driver fine
			fine, err := ph.policeService.GetDriverFine(ctx.Request().Context(), fineNum)
			if err != nil {
				logger.LogErr(err)
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}
			return ctx.JSON(http.StatusOK, fine) // response with a fine
		case "car":
			fine, err := ph.policeService.GetCarFine(ctx.Request().Context(), fineNum)
			if err != nil {
				logger.LogErr(err)
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}
			return ctx.JSON(http.StatusOK, fine) // response with a fine
		default:
			return ctx.JSON(http.StatusPreconditionFailed, "wrong fine type provided")
		}
	}
}

// RemoveFine removes driver/car fine with given fine num
func (ph *PoliceHandler) RemoveFine() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Get fine type from query parameter
		fineType := ctx.QueryParam("type")
		if fineType == "" {
			err := errors.New("fine type not specified")
			logger.LogErr(err)
			return ctx.JSON(http.StatusPreconditionRequired, err.Error())
		}
		// Get fine number from query parameter
		fineNum := ctx.QueryParam("num")
		if fineNum == "" {
			err := errors.New("fine number is empty")
			logger.LogErr(err)
			return ctx.JSON(http.StatusPreconditionRequired, err.Error())
		}

		// Choose fine type to delete from
		switch strings.ToLower(fineType) {
		case "driver": // Delete from driver fines
			err := ph.policeService.RemoveDriverFine(ctx.Request().Context(), fineNum)
			if err != nil {
				logger.LogErr(err)
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}

		case "car": // Delete from car fines
			err := ph.policeService.RemoveCarFine(ctx.Request().Context(), fineNum)
			if err != nil {
				logger.LogErr(err)
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}
		default:
			return ctx.JSON(http.StatusPreconditionFailed, "wrong fine type provided")
		}
		return ctx.JSON(http.StatusOK, "Fine successfully removed")
	}
}
