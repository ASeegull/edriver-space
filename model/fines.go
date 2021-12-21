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
	Cost      uint   `xml:"cost" json:"cost"`
}
