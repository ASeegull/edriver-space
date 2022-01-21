package model

type Car struct {
	ID                    string `json:"id"`
	Make                  string `json:"make"`
	Type                  string `json:"type"`
	CommercialDescription string `json:"commercial_description"`
	VIN                   string `json:"VIN_code"`
	MaxMass               int    `json:"maximum_mass"`
	ServiceMass           int    `json:"mass_of_the_vehicle_in_service"`
	VehicleCategory       string `json:"vehicle_category"`
	Capacity              int    `json:"capacity"`
	Colour                string `json:"colour_of_the_vehicle"`
	SeatsNum              int    `json:"number_of_seats_including_drivers_seat"`
	RegistrationNum       string `json:"registration_number"`
	FirstRegDate          string `json:"date_of_first_registration"`
	FullName              string `json:"full_name"`
	Address               string `json:"address"`
	Ownership             string `json:"ownership"`
	ValidityPeriod        string `json:"period_of_validity"`
	RegistrationDate      string `json:"date_of_registration"`
}
