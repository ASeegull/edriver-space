package server

import (
	"github.com/ASeegull/edriver-space/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// createUser creates new user
func (s *Server) createUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var u model.User // Store user info

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

// getUsers returns all users
func (s *Server) getUsers() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		users := make([]model.User, 0) // Store all users data

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

// getUser returns user by id
func (s *Server) getUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var u model.User                         // Store user info
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

// updateUser changes provided fields of user with given id
func (s *Server) updateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var u model.User                     // store user data to replace for
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

// deleteUser removes user
func (s *Server) deleteUser() echo.HandlerFunc {
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
