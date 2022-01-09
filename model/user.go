package model

type User struct {
	Id                  string
	Firstname           *string
	Lastname            *string
	Email               *string
	Password            *string
	Role                *string
	DriverLicenseNumber string
	Cars                []Car
}
