package model

import "encoding/xml"

type Data struct {
	XMLName      xml.Name      `xml:"data" json:"-"`
	ParkingFines []ParkingFine `xml:"parkingFine" json:"parking_fines"`
}

type ParkingFine struct {
	ID        string `json:"id"`
	FineNum   string `xml:"fineNum" json:"fine_num"`
	IssueTime string `xml:"issueTime" json:"issue_time"`
	CarVIN    string `xml:"carVIN" json:"car_VIN"`
	Cost      int    `xml:"cost" json:"cost"`
	PhotoURL  string `xml:"photo_url" json:"photo_url"`
}

func MakeParkingFine(FineNum, IssueTime, CarID string, Cost int, PhotoURL string) ParkingFine {
	return ParkingFine{FineNum: FineNum, IssueTime: IssueTime, CarVIN: CarID, Cost: Cost, PhotoURL: PhotoURL}
}
