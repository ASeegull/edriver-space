package server

import (
	"bytes"
	"errors"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/model"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

// uploadXMLFines writes parking fines data from xml file to the server(database in the future)
func (s *Server) uploadXMLFines() echo.HandlerFunc {
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

// uploadExcelFines parses uploaded Excel file with fines to the server(database in the future)
func (s *Server) uploadExcelFines() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// 1) Parse input(file), type multipart/form-data
		err := ctx.Request().ParseMultipartForm(10 << 20) // 10 Mb file maximum
		if err != nil {
			logger.LogErr(err)
			return ctx.String(http.StatusInternalServerError, "Error parsing input file")
		}

		// 2) Retrieve file from posted form-data
		FileKey := "File"
		file, header, err := ctx.Request().FormFile(FileKey)
		if err != nil {
			logger.LogErr(err)
			if header.Header.Get("name") != FileKey {
				return ctx.String(http.StatusInternalServerError, "Wrong key provided")
			}
			return ctx.String(http.StatusInternalServerError, "Error retrieving file from parsed data")
		}
		// Closing the file
		defer func(file multipart.File) {
			err = file.Close()
			if err != nil {
				logger.LogErr(err)
				panic(err)
			}
		}(file)

		// 3) Check if it is Excel (type XLSX/XLSM/XLTM/XLTX) file
		fn := strings.Split(header.Filename, ".") // Separate file name
		fileType := fn[len(fn)-1]

		// Check type
		ok := false
		if fileType == "xlsx" || fileType == "xlsm" || fileType == "xltm" || fileType == "xltx" {
			ok = true
		}
		if !ok {
			err = errors.New("wrong file type")
			logger.LogErr(err)
			return ctx.String(http.StatusNotAcceptable, err.Error())
		}

		// 4) Read from uploaded file
		// Create a buffer with a size of the uploaded file
		buf := make([]byte, header.Size)

		// Read from uploaded file into buffer
		_, err = file.Read(buf)
		if err != nil {
			logger.LogErr(err)
			return ctx.String(http.StatusInternalServerError, "Error reading from file")
		}

		// Create reader
		r := bytes.NewReader(buf)

		// 5) Retrieve data from the file
		// Open reader
		excelFile, err := excelize.OpenReader(r)
		if err != nil {
			logger.LogErr(err)
			return ctx.String(http.StatusInternalServerError, "Error opening reader")
		}

		// Keep retrieved parking fines here
		parkingFines := make([]model.ParkingFine, 0)

		// Indexes for getting data from rows
		const (
			IDCol = iota // 0
			IssueTimeCol
			CarIDCol
			CostCol
		)

		// Go through all sheets and collect all data
		for i := 0; i < excelFile.SheetCount; i++ {
			rows, err := excelFile.GetRows(excelFile.GetSheetName(i))
			if err != nil {
				logger.LogErr(err)
				return ctx.String(http.StatusInternalServerError, "Error retrieving rows from file.")
			}

			// Go through all rows
			for _, row := range rows {
				// Skip the first row with designation info
				if row[0] == "ID" || row[0] == "Id" || row[0] == "id" || row[0] == "iD" {
					continue
				}
				// Convert fine cost from string to int
				cost, err := strconv.Atoi(row[CostCol])
				if err != nil {
					logger.LogErr(err)
					return ctx.String(http.StatusInternalServerError, "Error converting string to int.")
				}
				// Create new parking fine
				parkingFine := model.MakeParkingFine(row[IDCol], row[IssueTimeCol], row[CarIDCol], cost)

				// Add parking fine to the slice
				parkingFines = append(parkingFines, parkingFine)
			}
		}

		/*
			Work with database, add all parkingFines values to the fines table
		*/

		// Only without database implementation
		s.Data.ParkingFines = append(s.Data.ParkingFines, parkingFines...)

		return ctx.String(http.StatusOK, "All fines successfully added.")
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
