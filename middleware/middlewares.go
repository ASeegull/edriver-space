package middleware

import (
	"errors"
	"github.com/ASeegull/edriver-space/logger"
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
	JWTAuthorization(roleWithAccess string) echo.MiddlewareFunc
}

type Middlewares struct {
	jwtManager auth.TokenManager
}

func NewMiddlewares(jwtManager auth.TokenManager) *Middlewares {
	return &Middlewares{jwtManager: jwtManager}
}

// UserIdentity checks if a user signed in and if yes, set the user's id in the context
func (mw *Middlewares) UserIdentity() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			userClaims, err := mw.parseAuthHeader(ctx)
			if err != nil {
				logger.LogErr(err)
				return ctx.JSON(http.StatusUnauthorized, err.Error())
			}

			ctx.Set(userCtx, userClaims.Id)
			return next(ctx)
		}
	}
}

// JWTAuthorization gives access to user if his role in access token is the same as role given as argument
func (mw *Middlewares) JWTAuthorization(roleWithAccess string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			userClaims, err := mw.parseAuthHeader(ctx)
			if err != nil {
				logger.LogErr(err)
				return ctx.String(http.StatusForbidden, "Unable to retrieve user role")
			}

			// Check whether to give access
			if userClaims.Role == roleWithAccess {
				return next(ctx) // Give access
			}
			return ctx.String(http.StatusForbidden, "Access denied!")
		}
	}
}

func (mw *Middlewares) parseAuthHeader(c echo.Context) (auth.UserClaims, error) {
	header := c.Request().Header.Get(authorizationHeader)
	if header == "" {
		return auth.UserClaims{}, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return auth.UserClaims{}, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return auth.UserClaims{}, errors.New("token is empty")
	}

	return mw.jwtManager.Parse(headerParts[1])
}
