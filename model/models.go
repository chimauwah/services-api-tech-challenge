package model

import (
	"time"
)

// An Employee model
//
// This is used for operations that involve an Employee
// swagger:parameters
type Employee struct {
	ID                    int       `json:"id"`
	EnterTs               time.Time `json:"enterts"`
	Active                bool      `json:"active"`
	FirstName             string    `json:"firstname"`
	LastName              string    `json:"lastname"`
	CellPhone             string    `json:"cellphone"`
	Title                 string    `json:"title"`
	Samaccountname        string    `json:"samaccountname"`
	Mail                  string    `json:"mail"`
	PrimaryPa             string    `json:"primarypa"`
	Office                string    `json:"office"`
	ManagerDn             string    `json:"managerdn"`
	ManagerSamaccountname string    `json:"managersamaccountname"`
	TravelPref            string    `json:"travelpref"`
}

// CoreSkill struct
type CoreSkill struct {
	Skill       string `json:"skill"`
	Proficiency string `json:"proficiency"`
}

// An EmployeeDetail model
//
// This is used for operations that involve an EmployeeDetail.
//
// swagger:parameters
type EmployeeDetail struct {
	Employee Employee    `json:"employee"`
	Skills   []CoreSkill `json:"skills"`
}
