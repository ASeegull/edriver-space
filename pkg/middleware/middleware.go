package middleware

import (
	"github.com/labstack/echo/v4"
)

// CustomMiddlewares stores all custom middleware
type CustomMiddlewares struct {
	Authorize Authorizer
}

// NewCustomMiddlewares returns a pointer to new CustomMiddlewares
func NewCustomMiddlewares() *CustomMiddlewares {
	return &CustomMiddlewares{Authorize: NewAuthorize()}
}

// Authorizer provides logic for user authorization
type Authorizer interface {
	JWTAuthorization(roleWithAccess string) echo.MiddlewareFunc
}
