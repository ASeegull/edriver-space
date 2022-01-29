package model

type Car struct {
	ID               string `json:"id"`
	Mark             string `json:"mark"`
	Type             string `json:"type"`
	VIN              string `json:"VIN_code"`
	MaxMass          int    `json:"maximum_mass"`
	VehicleCategory  string `json:"vehicle_category"`
	Colour           string `json:"colour_of_the_vehicle"`
	SeatsNum         int    `json:"number_of_seats_including_drivers_seat"`
	RegistrationNum  string `json:"registration_number"`
	FullName         string `json:"full_name"`
	ValidityPeriod   string `json:"period_of_validity"`
	RegistrationDate string `json:"date_of_registration"`
}
