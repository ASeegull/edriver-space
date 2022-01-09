package model

// Car fields may be not final
type Car struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	VIN             string `json:"VIN_code"`
	RegistrationNum string `json:"registration_number"`

	VehicleCategory       string `json:"vehicle_category"`
	Make                  string `json:"make"`
	Type                  string `json:"type"`
	CommercialDescription string `json:"commercial_description"`

	MaxMass     int    `json:"maximum_mass"`
	ServiceMass int    `json:"mass_of_the_vehicle_in_service"`
	Capacity    int    `json:"capacity"`
	Colour      string `json:"colour_of_the_vehicle"`
	SeatsNum    int    `json:"number_of_seats_including_drivers_seat"`

	FirstRegDate     string `json:"date_of_first_registration"`
	ValidityPeriod   string `json:"period_of_validity"`
	RegistrationDate string `json:"date_of_registration"`

	FullName  string `json:"full_name"`
	Address   string `json:"address"`
	Ownership string `json:"ownership"`
}
