package model

import (
	"github.com/ASeegull/edriver-space/api/server"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// User fields may be not final
type User struct {
	ID                  int    `json:"id"`
	LicenseNumber       string `json:"license_number"`
	DateOfIssue         string `json:"date_of_issue"`
	ExpireDate          string `json:"expire_date"`
	IndividualTaxNumber string `json:"individual_tax_number"`

	Category            string `json:"category"`
	CategoryIssuingDate string `json:"category_issuing_date"`
	CategoryExpire      string `json:"category_expire"`

	FullName     string `json:"full_name"`
	DateOfBirth  string `json:"date_of_birth"`
	PlaceOfBirth string `json:"place_of_birth"`

	Restrictions string `json:"restrictions"`
}

// CreateUser creates new user
func CreateUser(s *server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var u User // Store user info

		if err := ctx.Bind(&u); err != nil { // Binds http request data to provided argument
			return err
		}
		s.Users = append(s.Users, u) // add new user to the server

		/*
			Work with database, add new user with u fields
		*/

		return ctx.JSON(http.StatusCreated, u) // response with created user data
	}
}

// GetUsers returns all users
func GetUsers(s *server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		users := make([]User, 0) // Store all users data

		/*
			Work with database, assign values from database to users
		*/

		/*
			s.Users = users // this should only be used after database implementation
		*/
		_ = users                               // Temporary
		return ctx.JSON(http.StatusOK, s.Users) // response with all users data
	}
}

// GetUser returns user by id
func GetUser(s *server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var u User                               // Store user info
		ID, err := strconv.Atoi(ctx.Param("id")) // This will get an id parameter from the url
		if err != nil {
			return err
		}

		/*
			Work with database, get user with id == ID
		*/

		// Code below may be changed after database implementation
		for _, user := range s.Users { // Find user with provided id
			if user.ID == ID {
				u = user // assign found user to the returned user
			}
		}
		if u.ID == 0 {
			return ctx.JSON(http.StatusNotFound, "No user with such id")
		}
		return ctx.JSON(http.StatusOK, u) // response with found user data
	}
}

// UpdateUser changes provided fields of user with given id
func UpdateUser(s *server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var u User                           // store user data to replace for
		if err := ctx.Bind(&u); err != nil { // Binds http request data to provided argument
			return err
		}

		ID, err := strconv.Atoi(ctx.Param("id")) // This will get an id parameter from the url
		if err != nil {
			return err
		}

		/*
			Work with database, assign not empty values of u to fields of user with id == ID
		*/

		// Code below may be changed after database implementation
		for i, user := range s.Users { // find user with provided id
			if user.ID == ID { // update user data with data from request
				if u.LicenseNumber != "" {
					s.Users[i].LicenseNumber = u.LicenseNumber
				}
				if u.DateOfIssue != "" {
					s.Users[i].DateOfIssue = u.DateOfIssue
				}
				if u.ExpireDate != "" {
					s.Users[i].ExpireDate = u.ExpireDate
				}
				if u.IndividualTaxNumber != "" {
					s.Users[i].IndividualTaxNumber = u.IndividualTaxNumber
				}

				if u.Category != "" {
					s.Users[i].Category = u.Category
				}
				if u.CategoryIssuingDate != "" {
					s.Users[i].CategoryIssuingDate = u.CategoryIssuingDate
				}
				if u.CategoryExpire != "" {
					s.Users[i].CategoryExpire = u.CategoryExpire
				}

				if u.FullName != "" {
					s.Users[i].FullName = u.FullName
				}
				if u.DateOfBirth != "" {
					s.Users[i].DateOfBirth = u.DateOfBirth
				}
				if u.PlaceOfBirth != "" {
					s.Users[i].PlaceOfBirth = u.PlaceOfBirth
				}
				if u.Restrictions != "" {
					s.Users[i].Restrictions = u.Restrictions
				}
				u = s.Users[i] // swap needed to return not empty fields
			}
		}
		return ctx.JSON(http.StatusOK, u) // response with changed user data
	}
}

// DeleteUser removes user
func DeleteUser(s *server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ID, err := strconv.Atoi(ctx.Param("id")) // This will get an id parameter from the url
		if err != nil {
			return err
		}

		/*
			Work with database, delete user with id == ID
		*/

		// Code below may be changed after database implementation
		for i, user := range s.Users { // find user with provided id
			if user.ID == ID {
				s.Users = append(s.Users[:i], s.Users[i+1:]...) // remove user from server
			}
		}
		return ctx.String(http.StatusOK, "User successfully deleted.") // response with success message
	}
}
