package main

import "time"

// Device struct
type Device struct {
	ID             int       `json:"id"`
	EnterTs        time.Time `json:"enterts"`
	LastChangeTs   time.Time `json:"lastchangets"`
	Active         bool      `json:"active"`
	DeviceType     string    `json:"devicetype"`
	NotificationID string    `json:"notifcationid"`
	DeviceID       string    `json:"deviceId"`
	EmployeeID     int       `json:"employeeid"`
}
