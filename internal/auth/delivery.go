package auth

import "github.com/labstack/echo/v4"

type Handlers interface {
	Login() echo.HandlerFunc
	Logout() echo.HandlerFunc
	GetMe() echo.HandlerFunc
	GetCSRFToken() echo.HandlerFunc
}
