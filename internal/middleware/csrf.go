package middleware

import (
	"github.com/ASeegull/edriver-space/pkg/csrf"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (mw Middleware) CSRF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get(csrf.CSRFHeader)

		if token == "" {
			return c.JSON(http.StatusForbidden, "No CSRF Token")
		}

		sid, ok := c.Get("sid").(string)

		if !ok || !csrf.ValidateToken(token, sid) {
			return c.JSON(http.StatusForbidden, "Invalid CSRF Token")
		}

		return next(c)
	}
}
