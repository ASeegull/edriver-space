package model

import "encoding/xml"

type Data struct {
	XMLName      xml.Name      `xml:"data" json:"-"`
	ParkingFines []ParkingFine `xml:"parkingFine" json:"parking_fines"`
}

type ParkingFine struct {
	ID        string `xml:"fineID" json:"fine_id"`
	IssueTime string `xml:"issueTime" json:"issue_time"`
	CarID     string `xml:"carID" json:"car_id"`
	Cost      int    `xml:"cost" json:"cost"`
	PhotoURL  string `xml:"photo_url" json:"photo_url"`
}

func MakeParkingFine(ID, IssueTime, CarID string, Cost int, PhotoURL string) ParkingFine {
	return ParkingFine{ID: ID, IssueTime: IssueTime, CarID: CarID, Cost: Cost, PhotoURL: PhotoURL}
}
