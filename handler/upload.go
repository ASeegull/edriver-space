package handler

import (
	"bytes"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/middleware"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"strings"
)

type UploadHandler struct {
	UploadService service.Uploader
	cfg           *config.Config
}

// NewUploadHandler returns pointer to a new UploadHandler
func NewUploadHandler(UploadService service.Uploader, cfg *config.Config) *UploadHandler {
	return &UploadHandler{
		UploadService: UploadService,
		cfg:           cfg,
	}
}

func (u *UploadHandler) InitUploaderRoutes(e *echo.Group, mw middleware.Middleware) {
	// Upload routes (secure access)
	upload := e.Group("/upload", mw.JWTAuthorization("police"))

	upload.POST("/XML", u.UploadXMLFines()) // Upload XML fines data to the server
	upload.POST("/Excel", u.UploadExcel())
}

// UploadXMLFines reads fines data from xml and writes them to the server
func (u *UploadHandler) UploadXMLFines() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data model.Data // Data type stores slice of parking fines

		err := ctx.Bind(&data) // Bind slice of parking fines from the xml file
		if err != nil {
			logger.LogErr(err)
			return ctx.JSON(http.StatusBadRequest, "invalid input body")
		}

		err = u.UploadService.XMLFinesService(ctx.Request().Context(), data)
		if err != nil {
			logger.LogErr(err)
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, "Parking fines data successfully uploaded")
	}
}

// UploadExcel parses uploaded Excel file to the server and reads all fines from it
func (u *UploadHandler) UploadExcel() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Parse input(file), type multipart/form-data
		err := ctx.Request().ParseMultipartForm(10 << 20) // 10 Mb file maximum
		if err != nil {
			logger.LogErr(err)
			return ctx.JSON(http.StatusInternalServerError, "Error parsing input file")
		}

		// Retrieve file from posted form-data
		FileKey := "File"
		file, header, err := ctx.Request().FormFile(FileKey)
		if err != nil {
			logger.LogErr(err)
			if header.Header.Get("name") != FileKey {
				return ctx.JSON(http.StatusInternalServerError, "Wrong key provided")
			}
			return ctx.JSON(http.StatusInternalServerError, "Error retrieving file from parsed data")
		}
		// Closing the file
		defer func(file multipart.File) {
			err = file.Close()
			if err != nil {
				logger.LogErr(err)
				panic(err)
			}
		}(file)

		// Check if it is Excel (type XLSX/XLSM/XLTM/XLTX) file
		fn := strings.Split(header.Filename, ".") // Separate file name
		fileType := fn[len(fn)-1]

		// Check type
		ok := false
		if fileType == "xlsx" || fileType == "xlsm" || fileType == "xltm" || fileType == "xltx" {
			ok = true
		}
		if !ok {
			logger.LogErr(model.ErrWrongFileType)
			return ctx.JSON(http.StatusNotAcceptable, model.ErrWrongFileType.Error())
		}

		// Read from uploaded file into buffer
		// Create a buffer with a size of the uploaded file
		buf := make([]byte, header.Size)

		// Read from uploaded file into buffer
		_, err = file.Read(buf)
		if err != nil {
			logger.LogErr(err)
			return ctx.JSON(http.StatusInternalServerError, "Error reading from file")
		}

		// Create reader
		r := bytes.NewReader(buf)

		err = u.UploadService.ReadFinesExcel(ctx.Request().Context(), r)
		if err != nil {
			logger.LogErr(err)
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, "All fines successfully added")
	}
}
