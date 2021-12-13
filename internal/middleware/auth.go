package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (mw *Middleware) AuthSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(mw.cfg.Cookie.Name)
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, "error no cookie with name session_id")
			}
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		sessionID := cookie.Value

		sess, err := mw.sessUC.GetSessionByID(c.Request().Context(), sessionID)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		user, err := mw.authUC.GetUserByID(c.Request().Context(), sess.UserId)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("sid", sessionID)
		c.Set("session", sess)
		c.Set("user", user)

		return next(c)
	}
}

func (mw *Middleware) AuthJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerHeader := c.Request().Header.Get("Authorization")

		headerParts := strings.Split(bearerHeader, " ")
		if len(headerParts) != 2 {
			return c.JSON(http.StatusUnauthorized, "bearer token has invalid syntax")
		}

		tokenString := headerParts[1]

		userId, err := mw.jwtManager.Parse(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		user, err := mw.authUC.GetUserByID(c.Request().Context(), userId)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("user", map[string]string{
			"id":    user.Id,
			"login": user.Login,
			"role":  user.Role,
		})

		return next(c)
	}
}

func (mw Middleware) RoleMiddleware() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return c.JSON(http.StatusOK, "role")
		}
	}
}
