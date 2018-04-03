package main

import "time"

// An Employee model
//
// This is used for operations that involve an Employee
// swagger:parameters
type Employee struct {
	ID                    int       `json:"id"`
	EnterTs               time.Time `json:"enterts"`
	LastChangeTs          time.Time `json:"lastchangets"`
	Active                bool      `json:"active"`
	FirstName             string    `json:"firstname"`
	LastName              string    `json:"lastname"`
	Address1              string    `json:"address1"`
	Address2              string    `json:"address2"`
	City                  string    `json:"city"`
	State                 string    `json:"state"`
	Zip                   string    `json:"zip"`
	CellPhone             string    `json:"cellphone"`
	HomePhone             string    `json:"homephone"`
	Picture               []byte    `json:"picture"`
	Title                 string    `json:"title"`
	Role                  int       `json:"role"`
	IPPhone               string    `json:"ipphone"`
	Samaccountname        string    `json:"samaccountname"`
	Mail                  string    `json:"mail"`
	PrimaryPa             string    `json:"primarypa"`
	SecondaryPa           string    `json:"secondarypa"`
	Office                string    `json:"office"`
	ManagerDn             string    `json:"managerdn"`
	TravelPref            string    `json:"travelpref"`
	ManagerSamaccountname string    `json:"managersamaccountname"`
	LastHash              int       `json:"lasthash"`
	ImageHash             string    `json:"imagehash"`
	NickName              string    `json:"nickname"`
	ClientLoc             string    `json:"clientloc"`
}
