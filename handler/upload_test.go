package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/service"
	mock_service "github.com/ASeegull/edriver-space/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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

	fmt.Println(string(data))
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
				ParkingFines: []model.ParkingFine{
					{FineNum: "123a", IssueTime: "17.12.21 16:23", CarVIN: "BC 2304 AB", Cost: 500, PhotoURL: "https://1"},
					{FineNum: "237b", IssueTime: "20.12.21 19:40", CarVIN: "KA 2343 DB", Cost: 237, PhotoURL: "https://2"},
				}},
			inputContext: context.Background(),
			mockBehavior: func(s *mock_service.MockUploader, ctx context.Context, data model.Data) {
				s.EXPECT().XMLFinesService(ctx, data).Return(nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: "\"Parking fines data successfully uploaded\"\n",
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
			w.Header().Set("Content-Type", "application/xml")
			req := httptest.NewRequest("POST", "/uploadXML", bytes.NewBuffer(testCase.inputBody))

			// Perform request
			e.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}
}
