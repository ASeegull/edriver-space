package http

import (
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/middleware"
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers, mw *middleware.Middleware) {
	authGroup.POST("/login", h.Login())
	authGroup.POST("/logout", h.Logout())

	authGroup.Use(mw.AuthSessionMiddleware)
	authGroup.GET("/token", h.GetCSRFToken())
	authGroup.GET("/me", h.GetMe(), mw.CSRF)
}
