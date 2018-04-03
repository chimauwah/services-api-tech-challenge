package main

import "time"

// Checkin struct
type Checkin struct {
	ID             int       `json:"id"`
	EnterTs        time.Time `json:"enterts"`
	LastChangeTs   time.Time `json:"lastchangets"`
	Active         bool      `json:"active"`
	GeoID          string    `json:"geoid"`
	LAT            float32   `json:"lat"`
	LNG            float32   `json:"lng"`
	Name           string    `json:"name"`
	Distance       float32   `json:"distance"`
	NotifRequested int       `json:"notifrequested"`
	NotiExpiration int       `json:"notifrequsted"`
	EmployeeID     int       `json:"employeeid"`
}
