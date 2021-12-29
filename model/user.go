package model

type User struct {
	Id                  string
	Email               *string
	Password            *string
	Role                *string
	DriverLicenseNumber string
	Cars                []Car
}
