package middleware

import (
	"errors"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"

	userCtx = "userId"
)

type Middleware interface {
	UserIdentity() echo.MiddlewareFunc
}

type Middlewares struct {
	jwtManager auth.TokenManager
}

func NewMiddlewares(jwtManager auth.TokenManager) *Middlewares {
	return &Middlewares{jwtManager: jwtManager}
}

func (mw *Middlewares) UserIdentity() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, err := mw.parseAuthHeader(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			c.Set(userCtx, id)
			return next(c)
		}
	}
}

func (mw *Middlewares) parseAuthHeader(c echo.Context) (string, error) {
	header := c.Request().Header.Get(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return mw.jwtManager.Parse(headerParts[1])
}
