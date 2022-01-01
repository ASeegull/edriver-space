package service

import (
	"bytes"
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/repository"
	"github.com/xuri/excelize/v2"
	"strconv"
)

// UploadService stores upload logic
type UploadService struct {
	UploadRepos repository.Uploader
	cfg         *config.Config
}

// NewUploadService returns a pointer to new UploadService
func NewUploadService(repos *repository.Repositories, cfg *config.Config) *UploadService {
	return &UploadService{UploadRepos: repos.Uploader, cfg: cfg}
}

// XMLFinesService goes through all fines in data and passes each to the database query
func (u *UploadService) XMLFinesService(ctx context.Context, data model.Data) error {
	for _, fine := range data.ParkingFines {
		err := u.UploadRepos.AddFine(ctx, fine) // Adding each fine to the database
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
		return err
	}

	// Indexes for getting data from rows
	const (
		IDCol = iota // 0
		IssueTimeCol
		CarIDCol
		CostCol
		URLCol
	)

	// Go through all sheets and collect all data
	for i := 0; i < excelFile.SheetCount; i++ {
		rows, err := excelFile.GetRows(excelFile.GetSheetName(i))
		if err != nil {
			return err
		}
		// Go through all rows
		for _, row := range rows {
			// Skip the first row with designation info
			if row[IDCol] == "ID" || row[IDCol] == "Id" || row[IDCol] == "id" || row[IDCol] == "iD" {
				continue
			}
			// Convert fine cost from string to int
			cost, err := strconv.Atoi(row[CostCol])
			if err != nil {
				return err
			}
			// Create new parking fine
			parkingFine := model.MakeParkingFine(row[IDCol], row[IssueTimeCol], row[CarIDCol], cost, row[URLCol])

			// Pass parking fine to the database query
			err = u.UploadRepos.AddFine(ctx, parkingFine)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
