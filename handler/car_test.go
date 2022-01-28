package handler

import (
	"bytes"
	"context"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"

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
	type mockBehavior func(s *mock_service.MockCars, car *model.Car)

	testTable := []struct {
		name         string
		requestBody  string
		serviceInput *model.Car
		mockBehavior mockBehavior
		statusCode   int
	}{{
		name: "Valid input",
		requestBody: `{
			"id":                                    "1",
			"mark":                                  "BMW",
			"type":                                  "E31",
			"VIN_code":                              "1234567986",
			"maximum_mass":                           184,
			"vehicle_category":                      "M",
			"colour_of_the_vehicle":                 "red",
			"number_of_seats_including_drivers_seat": 4,
			"registration_number":                   "BC345966",
			"full_name":                             "Ivan Ivanov",
			"period_of_validity":                    "2024-01-01",
			"date_of_registration":                  "2020-01-01"
			}`,
		serviceInput: &model.Car{
			ID:               "1",
			Mark:             "BMW",
			Type:             "E31",
			VIN:              "1234567986",
			MaxMass:          184,
			VehicleCategory:  "M",
			Colour:           "red",
			SeatsNum:         4,
			RegistrationNum:  "BC345966",
			FullName:         "Ivan Ivanov",
			ValidityPeriod:   "2024-01-01",
			RegistrationDate: "2020-01-01",
		},
		mockBehavior: func(r *mock_service.MockCars, car *model.Car) {
			r.EXPECT().CreateCar(context.Background(), car).Return(car, nil).Times(1)
		},
		statusCode: 201,
	},
		{
			name: "Invalid Input",
			requestBody: `{
				"id":                                    "",
			    "mark":                                  "",
			    "type":                                  "E31",
			    "VIN_code":                              "1234567986",
			    "maximum_mass":                           184,
			    "vehicle_category":                      "M",
			    "colour_of_the_vehicle":                 "red",
			    "number_of_seats_including_drivers_seat": 4,
			    "registration_number":                   "BC345966",
			    "full_name":                             "Ivan Ivanov",
			    "period_of_validity":                    "2024-01-01",
			    "date_of_registration":                  "2020-01-01"
			}`,
			serviceInput: &model.Car{
				ID:               "",
				Mark:             "",
				Type:             "E31",
				VIN:              "1234567986",
				MaxMass:          184,
				VehicleCategory:  "M",
				Colour:           "red",
				SeatsNum:         4,
				RegistrationNum:  "BC345966",
				FullName:         "Ivan Ivanov",
				ValidityPeriod:   "2024-01-01",
				RegistrationDate: "2020-01-01",
			},
			mockBehavior: func(r *mock_service.MockCars, car *model.Car) {
				r.EXPECT().CreateCar(context.Background(), car).Return(nil, errors.New("error")).Times(1)
			},
			statusCode: 400,
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
