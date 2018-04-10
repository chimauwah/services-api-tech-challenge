package model

import (
	"database/sql"

	"github.com/chimauwah/services-api-tech-challenge/helper"
)

// An EmployeeCreateEmployeeReq request model
//
// This is used for operations that require an Employee as body of the request
//
// swagger:parameters createEmployee
type EmployeeCreateEmployeeReq struct {
	// The Employee to create
	//
	// in: body
	// required: true
	Employee *Employee `json:"employee"`
}

// An EmployeeUpdateEmployeeReq request model
//
// This is used for operations that require an Employee as body of the request
//
// swagger:parameters updateEmployee
type EmployeeUpdateEmployeeReq struct {
	// The Employee date with which to overwrite existing Employee
	//
	// in: body
	// required: true
	Employee *Employee `json:"employee"`
}

// An EmployeeResponse response model
//
// This is used for returning a response with status code and list of retrieved employees as body
//
// swagger:response employeeResponse
type EmployeeResponse struct {
	// in: body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
		// Employee model
		Payload []*Employee `json:"employees"`
	}
}

// An Employee model
//
// This is used for operations that involve an Employee
// swagger:parameters
type Employee struct {
	ID                    int             `json:"id"`
	EnterTs               helper.NullTime `json:"enterts"`
	Active                bool            `json:"active"`
	FirstName             sql.NullString  `json:"firstname"`
	LastName              sql.NullString  `json:"lastname"`
	CellPhone             sql.NullString  `json:"cellphone"`
	Title                 sql.NullString  `json:"title"`
	Samaccountname        sql.NullString  `json:"samaccountname"`
	Mail                  sql.NullString  `json:"mail"`
	PrimaryPa             sql.NullString  `json:"primarypa"`
	Office                sql.NullString  `json:"office"`
	ManagerDn             sql.NullString  `json:"managerdn"`
	ManagerSamaccountname sql.NullString  `json:"managersamaccountname"`
	TravelPref            sql.NullString  `json:"travelpref"`
}

// CoreSkill struct
type CoreSkill struct {
	Skill       sql.NullString `json:"skill"`
	Proficiency sql.NullString `json:"proficiency"`
}

// An EmployeeDetail model
//
// This is used for operations that involve an EmployeeDetail
// swagger:parameters
type EmployeeDetail struct {
	Employee Employee    `json:"employee"`
	Skills   []CoreSkill `json:"skills"`
}

// Success response
// swagger:response ok
type swaggScsResp struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Code int `json:"code"`
	}
}

// Boolean response
// swagger:response bool
type swaggBoolResp struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Code int `json:"code"`
		// Boolean true/false
		Data bool `json:"data"`
	}
}

// Error Bad Request
// swagger:response badReq
type swaggErrBadReq struct {
	// in:body
	Body struct {
		// HTTP status code 400 - Status Bad Request
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Forbidden
// swagger:response forbidden
type swaggErrForbidden struct {
	// in:body
	Body struct {
		// HTTP status code 403 - Forbidden
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Not Found
// swagger:response notFound
type swaggErrNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 - Not Found
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Conflict
// swagger:response conflict
type swaggErrConflict struct {
	// in:body
	Body struct {
		// HTTP status code 409 - Conflict
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Interval Server
// swagger:response internal
type swaggErrInternal struct {
	// in:body
	Body struct {
		// HTTP status code 500 - Internal server error
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// A ValidationError is a swagger response to represent error
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}
