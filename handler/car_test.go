package handler

import (
	"bytes"
	"context"
	"net/http/httptest"
	"testing"

	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/pkg/validator"
	"github.com/ASeegull/edriver-space/service"
	mock_service "github.com/ASeegull/edriver-space/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCarsHandler_CreateCar(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCars, car model.Car)

	testTable := []struct {
		name         string
		requestBody  string
		serviceInput model.Car
		mockBehavior mockBehavior
		statusCode   int
	}{{
		name: "Valid input",
		requestBody: `{
			"id":                                    "4",
			"make":                                  "test",
			"type":                                  "test",
			"commercial_description":                "test",
			"VIN_code":                              "sss",
			"maximum_mass":                           1,
			"mass_of_the_vehicle_in_service":         1,
			"vehicle_category":                      "rrr",
			"capacity":                               1,
			"colour_of_the_vehicle":                 "red",
			"number_of_seats_including_drivers_seat": 4,
			"registration_number":                   "test",
			"date_of_first_registration":            "2020-01-01",
			"full_name":                             "TEST",
			"address":                               "TEST",
			"ownership":                             "TEST",
			"period_of_validity":                    "2020-01-01",
			"date_of_registration":                  "2020-01-01"
			}`,
		serviceInput: model.Car{
			ID:                    "4",
			Make:                  "test",
			Type:                  "test",
			CommercialDescription: "test",
			VIN:                   "sss",
			MaxMass:               1,
			ServiceMass:           1,
			VehicleCategory:       "rrr",
			Capacity:              1,
			Colour:                "red",
			SeatsNum:              4,
			RegistrationNum:       "test",
			FirstRegDate:          "2020-01-01",
			FullName:              "TEST",
			Address:               "TEST",
			Ownership:             "TEST",
			ValidityPeriod:        "2020-01-01",
			RegistrationDate:      "2020-01-01",
		},
		mockBehavior: func(r *mock_service.MockCars, car model.Car) {
			r.EXPECT().CreateCar(context.Background(), car).Return(car, nil)
		},
		statusCode: 200,
	},
		{
			name: "Invalid Input",
			requestBody: `{
			"id":                                    "",
			"make":                                  "",
			"type":                                  "",
			"commercial_description":                "test",
			"VIN_code":                              "sss",
			"maximum_mass":                           1,
			"mass_of_the_vehicle_in_service":         1,
			"vehicle_category":                      "rrr",
			"capacity":                               1,
			"colour_of_the_vehicle":                 "red",
			"number_of_seats_including_drivers_seat": 4,
			"registration_number":                   "test",
			"date_of_first_registration":            "2020-01-01",
			"full_name":                             "TEST",
			"address":                               "TEST",
			"ownership":                             "TEST",
			"period_of_validity":                    "2020-01-01",
			"date_of_registration":                  "2020-01-01"
			}`,
			mockBehavior: func(r *mock_service.MockCars, car model.Car) {},
			statusCode:   400,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			cfg := &config.Config{}

			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockCars(c)

			testCase.mockBehavior(s, testCase.serviceInput)

			services := &service.Services{Cars: s}
			handler := NewHandlers(services, cfg)

			// Init Endpoint
			e := echo.New()
			e.Validator = validator.NewValidationUtil()

			e.POST("/cars/", handler.Cars.CreateCar())

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/cars/", bytes.NewBufferString(testCase.requestBody))
			req.Header.Set("Content-type", "application/json")

			// Make Request
			e.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, testCase.statusCode)

		})
	}
}
