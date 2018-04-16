package model

import (
	"database/sql"

	"github.com/chimauwah/services-api-tech-challenge/helper"
)

// An Employee model
//
// This is used for operations that involve an Employee
// swagger:parameters
type Employee struct {
	ID        int             `json:"id"`
	EnterTs   helper.NullTime `json:"enterts"`
	Active    bool            `json:"active"`
	FirstName sql.NullString  `json:"firstname"`
	LastName  sql.NullString  `json:"lastname"`
	CellPhone sql.NullString  `json:"cellphone"`
	// example: Black Panther
	Title                 sql.NullString `json:"title"`
	Samaccountname        sql.NullString `json:"samaccountname"`
	Mail                  sql.NullString `json:"mail"`
	PrimaryPa             sql.NullString `json:"primarypa"`
	Office                sql.NullString `json:"office"`
	ManagerDn             sql.NullString `json:"managerdn"`
	ManagerSamaccountname sql.NullString `json:"managersamaccountname"`
	TravelPref            sql.NullString `json:"travelpref"`
}

// CoreSkill struct
type CoreSkill struct {
	Skill       sql.NullString `json:"skill"`
	Proficiency sql.NullString `json:"proficiency"`
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
