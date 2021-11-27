package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Car fields may be not final
type Car struct {
	ID              int    `json:"id"`
	VIN             string `json:"VIN_code"`
	RegistrationNum string `json:"registration_number"`

	VehicleCategory       string `json:"vehicle_category"`
	Make                  string `json:"make"`
	Type                  string `json:"type"`
	CommercialDescription string `json:"commercial_description"`

	MaxMass     int    `json:"maximum_mass"`
	ServiceMass int    `json:"mass_of_the_vehicle_in_service"`
	Capacity    int    `json:"capacity"`
	Colour      string `json:"colour_of_the_vehicle"`
	SeatsNum    int    `json:"number_of_seats_including_drivers_seat"`

	FirstRegDate     string `json:"date_of_first_registration"`
	ValidityPeriod   string `json:"period_of_validity"`
	RegistrationDate string `json:"date_of_registration"`

	FullName  string `json:"full_name"`
	Address   string `json:"address"`
	Ownership string `json:"ownership"`
}

// createCar creates new car
func (s *Server) createCar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var c Car // Store car info

		if err := ctx.Bind(&c); err != nil { // Binds http request data to provided argument
			return err
		}
		s.Cars = append(s.Cars, c) // add new car to the server

		/*
			Work with database, add new car with c data
		*/

		return ctx.JSON(http.StatusCreated, c) // response with created car data
	}
}

// getCars returns all cars
func (s *Server) getCars() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cars := make([]Car, 0) // Store all cars

		/*
			Work with database, assign values from database to cars
		*/

		/*
			s.Cars = cars // this should only be used after database implementation
		*/
		_ = cars                               // Temporary
		return ctx.JSON(http.StatusOK, s.Cars) // response with all cars data
	}
}

// getCar returns car by id
func (s *Server) getCar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var c Car                                // Store car info
		ID, err := strconv.Atoi(ctx.Param("id")) // This will get an id parameter from the url
		if err != nil {
			return err
		}

		/*
			Work with database, get car with id == ID
		*/

		// Code below may be changed after database implementation
		for _, car := range s.Cars { // Find car with provided id
			if car.ID == ID {
				c = car // assign found car to the returned car
			}
		}
		if c.ID == 0 {
			return ctx.JSON(http.StatusNotFound, "No car with such id")
		}
		return ctx.JSON(http.StatusOK, c) // response with found car data
	}
}

// updateCar changes provided fields of car with given id
func (s *Server) updateCar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var c Car                            // store car data to replace for
		if err := ctx.Bind(&c); err != nil { // Binds http request data to provided argument
			return err
		}

		ID, err := strconv.Atoi(ctx.Param("id")) // This will get an id parameter from the url
		if err != nil {
			return err
		}

		/*
			Work with database, assign not empty values of c to fields of car with id == ID
		*/

		// Code below may be changed after database implementation
		for i, user := range s.Users { // find car with provided id
			if user.ID == ID { // update car data with data from request
				if c.VIN != "" {
					s.Cars[i].VIN = c.VIN
				}
				if c.RegistrationNum != "" {
					s.Cars[i].RegistrationNum = c.RegistrationNum
				}

				if c.VehicleCategory != "" {
					s.Cars[i].VehicleCategory = c.VehicleCategory
				}
				if c.Make != "" {
					s.Cars[i].Make = c.Make
				}
				if c.Type != "" {
					s.Cars[i].Type = c.Type
				}
				if c.CommercialDescription != "" {
					s.Cars[i].CommercialDescription = c.CommercialDescription
				}

				if c.MaxMass != 0 {
					s.Cars[i].MaxMass = c.MaxMass
				}
				if c.ServiceMass != 0 {
					s.Cars[i].ServiceMass = c.ServiceMass
				}
				if c.Capacity != 0 {
					s.Cars[i].Capacity = c.Capacity
				}
				if c.Colour != "" {
					s.Cars[i].Colour = c.Colour
				}
				if c.SeatsNum != 0 {
					s.Cars[i].SeatsNum = c.SeatsNum
				}

				if c.FirstRegDate != "" {
					s.Cars[i].FirstRegDate = c.FirstRegDate
				}
				if c.ValidityPeriod != "" {
					s.Cars[i].ValidityPeriod = c.ValidityPeriod
				}
				if c.RegistrationDate != "" {
					s.Cars[i].RegistrationDate = c.RegistrationDate
				}

				if c.FullName != "" {
					s.Cars[i].FullName = c.FullName
				}
				if c.Address != "" {
					s.Cars[i].Address = c.Address
				}
				if c.Ownership != "" {
					s.Cars[i].Ownership = c.Ownership
				}
				c = s.Cars[i] // swap needed to return not empty fields
			}
		}
		return ctx.JSON(http.StatusOK, c) // response with changed car data
	}
}

// deleteCar removes car
func (s *Server) deleteCar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ID, err := strconv.Atoi(ctx.Param("id")) // This will get an id parameter from the url
		if err != nil {
			return err
		}

		/*
			Work with database, delete car with id == ID
		*/

		// Code below may be changed after database implementation
		for i, car := range s.Cars { // find car with provided id
			if car.ID == ID {
				s.Cars = append(s.Cars[:i], s.Cars[i+1:]...) // remove car from the server
			}
		}
		return ctx.String(http.StatusOK, "Car successfully deleted.") // response with success message
	}
}
