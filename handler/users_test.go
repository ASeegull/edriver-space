package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/pkg/validator"
	"github.com/ASeegull/edriver-space/service"
	mock_service "github.com/ASeegull/edriver-space/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUsersHandlers_SignUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUsers, input service.UserSignUpInput)

	tests := []struct {
		name         string
		requestBody  string
		serviceInput service.UserSignUpInput
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name: "ok",
			requestBody: `{
			"firstname": "tom",
			"lastname": "kin",
			"email": "tom_kin@gmail.com",
			"password": "12345678"
			}`,
			serviceInput: service.UserSignUpInput{
				Firstname: "tom",
				Lastname:  "kin",
				Email:     "tom_kin@gmail.com",
				Password:  "12345678",
			},
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignUpInput) {
				r.EXPECT().SignUp(context.Background(), input).Return(service.Tokens{}, nil).Times(1)
			},
			statusCode:   200,
			responseBody: fmt.Sprintf("%s\n", `{"accessToken":"","refreshToken":""}`),
		},
		{
			name: "missing firstname",
			requestBody: `{
			"firstname": "",
			"lastname": "kin",
			"email": "tom_kin@gmail.com",
			"password": "12345678"
			}`,
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignUpInput) {},
			statusCode:   400,
			responseBody: fmt.Sprintf("%q\n", "invalid input body"),
		},
		{
			name: "missing lastname",
			requestBody: `{
			"firstname": "tom",
			"lastname": "",
			"email": "tom_kin@gmail.com",
			"password": "12345678"
			}`,
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignUpInput) {},
			statusCode:   400,
			responseBody: fmt.Sprintf("%q\n", "invalid input body"),
		},
		{
			name: "missing email",
			requestBody: `{
			"firstname": "tom",
			"lastname": "kin",
			"email": "",
			"password": "12345678"
			}`,
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignUpInput) {},
			statusCode:   400,
			responseBody: fmt.Sprintf("%q\n", "invalid input body"),
		},
		{
			name: "missing password",
			requestBody: `{
			"firstname": "tom",
			"lastname": "kin",
			"email": "tom_kin@gmail.com",
			"password": ""
			}`,
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignUpInput) {},
			statusCode:   400,
			responseBody: fmt.Sprintf("%q\n", "invalid input body"),
		},
		{
			name: "password too short",
			requestBody: `{
			"firstname": "tom"
			"lastname": "kin",
			"email": "tom_kin@gmail.com",
			"password": "1234"
			}`,
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignUpInput) {},
			statusCode:   400,
			responseBody: fmt.Sprintf("%q\n", "input body has not json format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init Dependencies
			cfg := &config.Config{}

			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockUsers(c)

			tt.mockBehavior(s, tt.serviceInput)

			services := &service.Services{Users: s}
			handler := NewHandlers(services, cfg)

			// Init Endpoint
			e := echo.New()
			e.Validator = validator.NewValidationUtil()

			e.GET("/sign-up", handler.Users.SignUp())

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/sign-up", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-type", "application/json")

			// Make Request
			e.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}
}

func TestUsersHandlers_SignIn(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUsers, input service.UserSignInInput)

	tests := []struct {
		name         string
		requestBody  string
		serviceInput service.UserSignInInput
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name: "ok",
			requestBody: `
			{
			"email": "tom_kin@gmail.com",
			"password": "12345678"
			}`,
			serviceInput: service.UserSignInInput{
				Email:    "tom_kin@gmail.com",
				Password: "12345678",
			},
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignInInput) {
				r.EXPECT().SignIn(context.Background(), input).Return(service.Tokens{}, nil).Times(1)
			},
			statusCode:   200,
			responseBody: fmt.Sprintf("%s\n", `{"accessToken":"","refreshToken":""}`),
		},
		{
			name: "missing email",
			requestBody: `
			{
			"password": "12345678"
			}`,
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignInInput) {},
			statusCode:   400,
			responseBody: fmt.Sprintf("%q\n", "invalid input body"),
		},
		{
			name: "missing password",
			requestBody: `
			{
			"email": "tom_kin@gmail.com"
			}`,
			mockBehavior: func(r *mock_service.MockUsers, input service.UserSignInInput) {},
			statusCode:   400,
			responseBody: fmt.Sprintf("%q\n", "invalid input body"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init Dependencies
			cfg := &config.Config{}

			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockUsers(c)

			tt.mockBehavior(s, tt.serviceInput)

			services := &service.Services{Users: s}
			handler := NewHandlers(services, cfg)

			// Init Endpoint
			e := echo.New()
			e.Validator = validator.NewValidationUtil()

			e.GET("/sign-in", handler.Users.SignIn())

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/sign-in", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-type", "application/json")

			// Make Request
			e.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}
}

func TestUsersHandlers_AddDriverLicence(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUsers, input service.AddDriverLicenceInput)

	userId := "1"

	tests := []struct {
		name         string
		requestBody  string
		serviceInput service.AddDriverLicenceInput
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name: "ok",
			requestBody: `{"individual_tax_number": "1111222233334444"}`,
			serviceInput: service.AddDriverLicenceInput{IndividualTaxNumber: "1111222233334444"},
			mockBehavior: func(r *mock_service.MockUsers, input service.AddDriverLicenceInput) {
				r.EXPECT().AddDriverLicence(context.Background(), input, userId).Return(nil).Times(1)
			},
			statusCode: 200,
			responseBody: fmt.Sprintf("%q\n", "successfully added"),
		},
		{
			name: "missing individual tax number",
			requestBody: `{"individual_tax_number": ""}`,
			serviceInput: service.AddDriverLicenceInput{IndividualTaxNumber: "1111222233334444"},
			mockBehavior: func(r *mock_service.MockUsers, input service.AddDriverLicenceInput) {},
			statusCode: 400,
			responseBody: fmt.Sprintf("%q\n", "invalid input body"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init Dependencies
			cfg := &config.Config{}

			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockUsers(c)

			tt.mockBehavior(s, tt.serviceInput)

			services := &service.Services{Users: s}
			handler := NewHandlers(services, cfg)

			// Init Endpoint
			e := echo.New()
			e.Validator = validator.NewValidationUtil()

			e.GET("/add-driver-licence", handler.Users.AddDriverLicence(), func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					c.Set("userId", userId)
					return next(c)
				}
			})

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/add-driver-licence", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-type", "application/json")

			// Make Request
			e.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}
}
