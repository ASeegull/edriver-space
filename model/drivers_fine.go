package model

type DriversFine struct {
	Id                        string `json:"id"`
	LicenceNumber             string `json:"licence_number"`
	FineNum                   string `json:"fine_num"`
	DataAndTime               string `json:"data_and_time"`
	Place                     string `json:"place"`
	FileLawArticle            string `json:"file_law_article"`
	Price                     int    `json:"price"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
}
