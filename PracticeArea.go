package main

import "time"

// PracticeArea struct
type PracticeArea struct {
	ID           int       `json:"id"`
	EnterTs      time.Time `json:"enterts"`
	LastChangeTs time.Time `json:"lastchangets"`
	Active       bool      `json:"active"`
	Code         string    `json:"code"`
	Description  string    `json:"description"`
	PrimaryOnly  bool      `json:"primaryonly"`
	Sequence     int       `json:"sequence"`
}
