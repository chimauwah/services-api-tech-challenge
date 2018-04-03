package main

import "time"

// CoreSkill struct
type CoreSkill struct {
	ID           int       `json:"id"`
	EnterTs      time.Time `json:"enterts"`
	LastChangeTs time.Time `json:"lastchangets"`
	Active       bool      `json:"active"`
	Skill        string    `json:"skill"`
	Sequence     int       `json:"sequence"`
	EmployeeID   int       `json:"employeeid"`
	Proficiency  string    `json:"proficiency"`
}
