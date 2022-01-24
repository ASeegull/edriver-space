package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/repository"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

// UploadService stores upload logic
type UploadService struct {
	CarFinesRepos repository.CarFines
	cfg           *config.Config
}

// NewUploadService returns a pointer to new UploadService
func NewUploadService(repos *repository.Repositories, cfg *config.Config) *UploadService {
	return &UploadService{CarFinesRepos: repos.CarFines, cfg: cfg}
}

// XMLFinesService goes through all car fines in data and passes each to the database query
func (u *UploadService) XMLFinesService(ctx context.Context, data model.Data) error {
	for _, fine := range data.CarsFines {
		err := u.CarFinesRepos.AddCarFine(ctx, &fine) // Adding each fine to the database
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadFinesExcel reads all fines from Excel file and passes each to the database query
func (u *UploadService) ReadFinesExcel(ctx context.Context, r *bytes.Reader) error {
	// Retrieve data from the file
	// Open reader
	excelFile, err := excelize.OpenReader(r)
	if err != nil {
		err = errors.New("error opening Excel file")
		return err
	}

	// Indexes for getting data from rows
	const (
		regNumCol = iota // 0
		fineNumCol
		dateCol
		placeCol
		FLACol
		priceCol
		infoCol
		URLCol
	)

	// Go through all sheets and collect all data
	for i := 0; i < excelFile.SheetCount; i++ {
		rows, err := excelFile.GetRows(excelFile.GetSheetName(i))
		if err != nil {
			err = errors.New("error retrieving rows from the file")
			return err
		}
		// Go through all rows
		for _, row := range rows {
			// Skip the first row with designation info
			if strings.ToLower(row[regNumCol]) == "regnum" {
				continue
			}
			// Convert fine price from string to int
			price, err := strconv.Atoi(row[priceCol])
			if err != nil {
				err = errors.New("error converting string to int")
				return err
			}
			// Create new car fine
			carFine := model.NewCarsFine(row[regNumCol], row[fineNumCol], row[dateCol], row[placeCol], row[FLACol], price, row[infoCol], row[URLCol])

			// Pass car fine to the database query
			err = u.CarFinesRepos.AddCarFine(ctx, carFine)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
