package middleware

import (
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Authorize struct {
	*auth.Manager
}

// NewAuthorize return pointer to new Authorize
func NewAuthorize() *Authorize {
	return &Authorize{}
}

// JWTAuthorization gives access to user if his role in access token is the same as role given as argument
func (a *Authorize) JWTAuthorization(roleWithAccess string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Get authorization header as string
			bearerHeader := ctx.Request().Header.Get("Authorization")

			// Split header to get token parts
			token := strings.Split(bearerHeader, " ")
			if token[1] != "" {
				err := model.ErrJWTEmpty
				logger.LogErr(err)
				return ctx.String(http.StatusForbidden, "Unable to read from JWT token")
			}

			// Get claims with user role from the JWT token
			userRole, err := a.Parse(token[1])
			if err != nil {
				logger.LogErr(err)
				return ctx.String(http.StatusForbidden, "Unable to retrieve user role")
			}

			// Check whether to give access
			if userRole == roleWithAccess {
				return next(ctx) // Give access
			}
			return ctx.String(http.StatusForbidden, "Access denied!")
		}
	}
}
