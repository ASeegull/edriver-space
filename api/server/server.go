package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	*echo.Echo
}

// NewServer creates new server
func NewServer() *Server {
	s := &Server{
		Echo: echo.New(),
	}
	s.routes()
	return s
}

// routes stores all possible routes
func (s *Server) routes() {
	s.GET("/", hello(), middleware.Logger())             // Home
	s.GET("/version", getVersion(), middleware.Logger()) // Get project version
}

// getVersion returns current version of the app
func getVersion() echo.HandlerFunc {
	return func(context echo.Context) error {
		version := "0.0" // Use environment variable instead
		return context.String(http.StatusOK, version)
	}
}

//
func hello() echo.HandlerFunc {
	return func(context echo.Context) error {
		hi := "Hello Lv-644.Go!"
		return context.String(http.StatusOK, hi)
	}
}
