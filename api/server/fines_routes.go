package server

import (
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// uploadXMLData writes parking fines data from xml file to the server(database in the future)
func (s *Server) uploadXMLData() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data model.Data // Data type stores slice of parking fines

		err := ctx.Bind(&data) // Bind slice of parking fines from the xml file
		if err != nil {
			logger.LogErr(err)
			return ctx.String(http.StatusInternalServerError, "error binding data from xml file")
		}

		/*
			Work with database, add parking fines from data to the parking fines database
		*/

		s.Data.ParkingFines = append(s.Data.ParkingFines, data.ParkingFines...) // remove after database implementation

		return ctx.JSON(http.StatusOK, "Parking fines data was successfully uploaded")
	}
}

// getParkingFines responses with all parking fines from the server(database in future)
func (s *Server) getParkingFines() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		parkingFines := make([]model.ParkingFine, 0) // store parking fines from the database

		/*
			Work with database, assign all parking fines to the parkingFines
		*/

		_ = parkingFines // remove after database implementation
		return ctx.JSON(http.StatusOK, s.Data.ParkingFines)
	}
}

// getParkingFine responses with parking fine with provided in the url id
func (s *Server) getParkingFine() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var pf model.ParkingFine // store parking fine from the database

		ID := ctx.Param("id") // This will get an id parameter from the url
		if ID == "" {
			return ctx.JSON(http.StatusNotFound, "ID not specified")
		}

		/*
			Work with database, get parking fine with id == ID
		*/

		// Code below may be removed after database implementation
		for _, parkingFine := range s.Data.ParkingFines { // Find parking fine with provided id
			if parkingFine.ID == ID {
				pf = parkingFine // assign found parking fine to the returned parking fine
			}
		}

		return ctx.JSON(http.StatusFound, pf)
	}
}

// deleteParkingFine removes parking fine from the server(database in the future), by the id in the url
func (s *Server) deleteParkingFine() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ID := ctx.Param("id") // This will get an id parameter from the url
		if ID == "" {
			return ctx.JSON(http.StatusNotFound, "ID not specified")
		}

		/*
			Work with database, delete fine with id == fineID
		*/

		// Code below may be removed after database implementation
		for i, parkingFine := range s.Data.ParkingFines {
			if parkingFine.ID == ID {
				s.Data.ParkingFines = append(s.Data.ParkingFines[:i], s.Data.ParkingFines[i+1:]...) // remove parkingFine from server
			}
		}
		return ctx.JSON(http.StatusOK, "Parking fine successfully deleted")
	}
}
