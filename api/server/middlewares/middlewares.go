package middlewares

import (
	"errors"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// CustomMiddleware enables usage of methods of other types
type CustomMiddleware struct {
	*auth.Manager
}

// NewCustomMiddleware creates new CustomMiddleware
func NewCustomMiddleware() *CustomMiddleware {
	return &CustomMiddleware{}
}

// Authenticator provides logic for user authentication
type Authenticator interface {
	JWTAuthentication(role string) echo.MiddlewareFunc
}

// JWTAuthentication gives access to user if his role in access token is the same as role given as argument
func (cm *CustomMiddleware) JWTAuthentication(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Get authorization header as string
			bearerHeader := ctx.Request().Header.Get("Authorization")

			// Split header to get token parts
			token := strings.Split(bearerHeader, " ")
			if token[1] != "" {
				err := errors.New("JWT token has empty body")
				logger.LogErr(err)
				return ctx.String(http.StatusForbidden, "Unable to read from JWT token")
			}

			// Get claims with user role from the JWT token
			userRole, err := cm.Parse(token[1])
			if err != nil {
				logger.LogErr(err)
				return ctx.String(http.StatusForbidden, "Unable to retrieve user role")
			}

			// Check whether to give access
			if userRole == role {
				return next(ctx)
			}
			return ctx.String(http.StatusForbidden, "Access denied!")
		}
	}
}
