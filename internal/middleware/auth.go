package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (mw *Middleware) AuthSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(mw.cfg.Session.Name)
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

		user, err := mw.authUC.GetUserByID(c.Request().Context(), sess.UserID)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("sid", sessionID)
		c.Set("session", sess)
		c.Set("user", user)

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
