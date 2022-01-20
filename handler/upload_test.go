package handler

import (
	"bytes"
	"context"
	"encoding/xml"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/service"
	mock_service "github.com/ASeegull/edriver-space/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadHandler_UploadXMLFines(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUploader, ctx context.Context, data model.Data)

	data, err := os.ReadFile("../test/test_XML_upload.XML")
	if err != nil {
		logger.LogErr(err)
	}

	xmlNameData := xml.Name{Local: "data"}

	testTable := []struct {
		name                 string
		inputBody            []byte
		inputData            model.Data
		inputContext         context.Context
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{{
		name:         "Invalid input",
		inputBody:    []byte{1, 23, 34},
		inputData:    model.Data{},
		inputContext: context.Background(),
		mockBehavior: func(s *mock_service.MockUploader, ctx context.Context, data model.Data) {
			s.EXPECT().XMLFinesService(ctx, data).Return(nil).AnyTimes()
		},
		expectedStatusCode:   400,
		expectedResponseBody: "\"invalid input body\"\n",
	},
		{
			name:      "Valid input",
			inputBody: data,
			inputData: model.Data{
				XMLName: xmlNameData,
				CarsFines: []model.CarsFine{
					{VehicleRegistrationNumber: "RegNum1",
						FineNum: "fineNum1", DataAndTime: "20.01.22 18:30",
						Place: "Lviv", FileLawArticle: "fileLawArticle1",
						Price: 500, Info: "info1", ImdUrl: "https://1",
					},
					{VehicleRegistrationNumber: "RegNum2",
						FineNum: "fineNum2", DataAndTime: "22.01.22 18:30",
						Place: "Kyiv", FileLawArticle: "fileLawArticle2",
						Price: 700, Info: "info2", ImdUrl: "https://2",
					},
				}},
			inputContext: context.Background(),
			mockBehavior: func(s *mock_service.MockUploader, ctx context.Context, data model.Data) {
				s.EXPECT().XMLFinesService(ctx, data).Return(nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: "\"Fines data successfully uploaded\"\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init
			ctrl := gomock.NewController(t)

			cfg := &config.Config{}

			upload := mock_service.NewMockUploader(ctrl)
			testCase.mockBehavior(upload, testCase.inputContext, testCase.inputData)

			services := &service.Services{Uploader: upload}
			handler := NewHandlers(services, cfg)

			// Test server
			e := echo.New()
			e.POST("/uploadXML", handler.Upload.UploadXMLFines())

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/uploadXML", bytes.NewBuffer(testCase.inputBody))
			req.Header.Set("Content-Type", "application/xml;charset=UTF-8")

			// Perform request
			e.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}
}

func TestUploadHandler_UploadExcel(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUploader, ctx context.Context, r *bytes.Reader)

	wrongDataPath := "../test/test_XML_upload.XML"

	goodDataPath := "../test/test_Excel_upload_good.xlsx"

	wrongData, err := os.ReadFile(wrongDataPath)
	if err != nil {
		logger.LogErr(err)
	}

	goodData, err := os.ReadFile(goodDataPath)
	if err != nil {
		logger.LogErr(err)
	}

	fieldName := "File"

	testTable := []struct {
		name                 string
		inputBody            []byte
		filePath             string
		inputReader          *bytes.Reader
		inputContext         context.Context
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:         "Wrong file type",
			inputBody:    wrongData,
			filePath:     wrongDataPath,
			inputReader:  bytes.NewReader(wrongData),
			inputContext: context.Background(),
			mockBehavior: func(s *mock_service.MockUploader, ctx context.Context, r *bytes.Reader) {
				s.EXPECT().ReadFinesExcel(ctx, r).Return(nil).AnyTimes()
			},
			expectedStatusCode:   406,
			expectedResponseBody: "\"wrong file type\"\n",
		},
		{
			name:         "Good file type",
			inputBody:    goodData,
			filePath:     goodDataPath,
			inputReader:  bytes.NewReader(goodData),
			inputContext: context.Background(),
			mockBehavior: func(s *mock_service.MockUploader, ctx context.Context, r *bytes.Reader) {
				s.EXPECT().ReadFinesExcel(ctx, r).Return(nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: "\"All fines successfully added\"\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init
			ctrl := gomock.NewController(t)

			cfg := &config.Config{}

			upload := mock_service.NewMockUploader(ctrl)
			testCase.mockBehavior(upload, testCase.inputContext, testCase.inputReader)

			services := &service.Services{Uploader: upload}
			handler := NewHandlers(services, cfg)

			body := new(bytes.Buffer)

			mw := multipart.NewWriter(body)

			file, _ := os.Open(testCase.filePath)
			if err != nil {
				t.Fatal(err)
			}

			w, _ := mw.CreateFormFile(fieldName, testCase.filePath)
			if err != nil {
				t.Fatal(err)
			}

			if _, err := io.Copy(w, file); err != nil {
				t.Fatal(err)
			}

			// close the writer before making the request
			mw.Close()

			// Test server
			e := echo.New()
			e.POST("/uploadExcel", handler.Upload.UploadExcel())

			// Test request
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/uploadExcel", body)

			req.Header.Add("Content-Type", mw.FormDataContentType())

			// Perform request
			e.ServeHTTP(rec, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, rec.Code)
			assert.Equal(t, testCase.expectedResponseBody, rec.Body.String())

		})
	}
}
