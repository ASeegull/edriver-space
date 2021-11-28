package server

import (
	"github.com/ASeegull/edriver-space/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	*echo.Echo
	Users []model.User
	Cars  []model.Car
}

// NewServer starts new Server
func NewServer() *Server {
	s := &Server{
		Echo:  echo.New(),
		Users: []model.User{},
		Cars:  []model.Car{},
	}
	// All possible routes
	s.routes()
	return s
}

// routes stores all possible routes
func (s *Server) routes() {
	s.GET("/", s.hello())             // Home
	s.GET("/Version", s.getVersion()) // Get project version

	// Users CRUD
	s.GET("/Users", model.GetUsers(s))         // Get all users
	s.GET("/User/:id", model.GetUser(s))       // Get user by id
	s.POST("/User", model.CreateUser(s))       // Create new user
	s.PATCH("/User/:id", model.UpdateUser(s))  // Update user data
	s.DELETE("/User/:id", model.DeleteUser(s)) // Delete user

	// Cars CRUD
	s.GET("/Cars", model.GetCars(s))         // Get all cars
	s.GET("/Car/:id", model.GetCar(s))       // Get car by id
	s.POST("/Car", model.CreateCar(s))       // Create new car
	s.PATCH("/Car/:id", model.UpdateCar(s))  // Update car data
	s.DELETE("/Car/:id", model.DeleteCar(s)) // Delete car
}

// getVersion returns current version of the app
func (s *Server) getVersion() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		version := "0.0" // Use environment variable instead
		return ctx.JSON(http.StatusOK, version)
	}
}

func (s *Server) hello() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hi := "Hello Lv-644.Go!"
		return ctx.JSON(http.StatusOK, hi)
	}
}
