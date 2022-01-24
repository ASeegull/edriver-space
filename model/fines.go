package model

import "encoding/xml"

type Fines struct {
	DriversFines []DriversFine `json:"drivers_fines"`
	CarsFines    []CarsFine    `json:"cars_fines"`
}

type Data struct {
	XMLName   xml.Name   `xml:"data" json:"-"`
	CarsFines []CarsFine `xml:"carsFine" json:"cars_fines"`
}

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

type CarsFine struct {
	Id                        string `json:"id" xml:"id"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number" xml:"regNum"`
	FineNum                   string `json:"fine_num" xml:"fineNum"`
	DataAndTime               string `json:"data_and_time" xml:"dataAndTime"`
	Place                     string `json:"place" xml:"place"`
	FileLawArticle            string `json:"file_law_article" xml:"fileLawArticle"`
	Price                     int    `json:"price" xml:"price"`
	Info                      string `json:"info" xml:"info"`
	ImdUrl                    string `json:"imd_url" xml:"imdURL"`
}

// NewDriversFine returns a pointer to a new DriversFine
func NewDriversFine(licenceNumber, fineNum, dataTime, place, fileLawArticle string, price int, regNum string) *DriversFine {
	return &DriversFine{
		LicenceNumber:             licenceNumber,
		FineNum:                   fineNum,
		DataAndTime:               dataTime,
		Place:                     place,
		FileLawArticle:            fileLawArticle,
		Price:                     price,
		VehicleRegistrationNumber: regNum,
	}
}

// NewCarsFine returns a pointer to a new CarsFine
func NewCarsFine(regNum, fineNum, dataTime, place, fileLawArticle string, price int, info, imdURL string) *CarsFine {
	return &CarsFine{
		VehicleRegistrationNumber: regNum,
		FineNum:                   fineNum,
		DataAndTime:               dataTime,
		Place:                     place,
		FileLawArticle:            fileLawArticle,
		Price:                     price,
		Info:                      info,
		ImdUrl:                    imdURL,
	}
}
