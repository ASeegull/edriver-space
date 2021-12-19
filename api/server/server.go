package server

import (
	"github.com/ASeegull/edriver-space/api/server/middlewares"
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

	// Create custom middleware for authentication
	var auth middlewares.Authenticator = middlewares.NewCustomMiddleware()

	user := s.Group("/user", auth.JWTAuthentication("police"))
	car := s.Group("/car", auth.JWTAuthentication("ParkingManager"))

	// Users CRUD
	user.GET("", s.getUsers())          // Get all users
	user.GET("/:id", s.getUser())       // Get user by id
	user.POST("", s.createUser())       // Create new user
	user.PATCH("/:id", s.updateUser())  // Update user data
	user.DELETE("/:id", s.deleteUser()) // Delete user

	// Cars CRUD
	car.GET("", s.getCars())          // Get all cars
	car.GET("/:id", s.getCar())       // Get car by id
	car.POST("", s.createCar())       // Create new car
	car.PATCH("/:id", s.updateCar())  // Update car data
	car.DELETE("/:id", s.deleteCar()) // Delete car
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
