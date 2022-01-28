package model

// Driver fields may be not final
type Driver struct {
	ID                  string `json:"id"`
	FullName            string `json:"full_name"`
	DateOfBirth         string `json:"date_of_birth"`
	PlaceOfBirth        string `json:"place_of_birth"`
	DateOfIssue         string `json:"date_of_issue"`
	ExpireDate          string `json:"expire_date"`
	LicenseNumber       string `json:"license_number"`
	Category            string `json:"category"`
	CategoryIssuingDate string `json:"category_issuing_date"`
	IndividualTaxNumber string `json:"individual_tax_number"`
}
