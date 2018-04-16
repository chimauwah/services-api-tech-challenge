package model

import (
	"database/sql"
)

// A GetEmployeeByIDPathParam model
//
// This is used for path parameters for get employee by id.
//
// swagger:parameters getEmployee
type swaggGetEmployeeByIDParam struct {
	// id of employee to retrieve
	// in: path
	// required: true
	// example: 1710
	ID int32 `json:"id"`
}

// An UpdateEmployeeReq request model
//
// This is used for updating an existing Employee.
//
// swagger:parameters updateEmployee
type swaggUpdateEmployeeReq struct {
	// The id of Employee to update.
	// in: path
	// required: true
	// example: 1710
	ID int32 `json:"id"`
	// The employee data with which to overwrite existing Employee.
	// in: body
	// required: true
	Employee Employee `json:"employee"`
}

// A DeleteEmployeePathParam model
//
// This is the path parameter for deleting an employee by id.
//
// swagger:parameters deleteEmployee
type swaggDeleteEmployeeParam struct {
	// id of employee to delete
	// in: path
	// required: true
	// example: 1710
	ID int32 `json:"id"`
}

// A PutPathParam model
//
// This is used for the path parameter for updating an employee.
//
// swagger:parameters updateEmployee
type swaggGetPathParam struct {
	// id of employee to retrieve
	// in: path
	// required: true
	// example: 1710
	ID int32 `json:"id"`
}

// A SearchQueryParam model
//
// This is used for query parameters for search.
//
// swagger:parameters searchEmployees
type swaggSearchQueryParams struct {
	// search by first name
	// in: query
	// example: James
	FirstName string `json:"first_name"`
	// search last name
	// in: query
	// example: Howlett
	LastName string `json:"last_name"`
	// search by title
	// in: query
	// example: Mutant
	Title string `json:"title"`
	// search by active (1 or 2)
	// in: query
	// pattern: [0-1]{1}
	// example: 1
	Active int32 `json:"active"`
	// search by cellphone
	// in: query
	// example: (555) 867-5309
	CellPhone string `json:"cell_phone"`
	// search by samaccountname
	// in: query
	// example: logan
	Samaccountname string `json:"samaccountname"`
	// search by email
	// in: query
	// example: logan@xmen.com
	Mail string `json:"mail"`
	// search by primary pa
	// in: query
	// example: OPS
	PrimaryPa string `json:"primary_pa"`
	// search by office
	// in: query
	// example: Alberta
	Office string `json:"office"`
	// search by manager samaccountname
	// in: query
	// example: professorx
	ManagerSamaccountname string `json:"manager_samaccountname"`
	// search by travel preference
	// in: query
	// example: Potentially
	TravelPref string `json:"travel_pref"`
}

// An EmployeeCreateEmployeeReq request model
//
// This is used for operations that require an Employee as body of the request.
//
// swagger:parameters createEmployee
type swaggCreateEmployeeReq struct {
	// The Employee to create.
	// in: body
	// required: true
	Employee Employee `json:"employee"`
}

// An EmployeeResponse response model
// swagger:response employeeResponse
type swaggEmployeeResponse struct {
	// Employee model
	Employee Employee `json:"employee"`
}

// An EmployeesResponse response model
// swagger:response employeesResponse
type swaggEmployeesResponse struct {
	// Employee model
	Employees Employee `json:"employees"`
}

// An EmployeeDetailResponse response model
// swagger:response employeeDtlResponse
type swaggEmployeeDtlResponse struct {
	// EmployeeDetail model
	EmployeeDetails EmployeeDetail `json:"employeeDetails"`
}

// Success
// swagger:response successResponse
type swaggSuccessResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		// example: 200
		Status int `json:"status"`
		// Request call success - true
		// example: true
		Success bool `json:"success"`
		// Detailed success message
		// example: Operation completed successfully
		Message string `json:"message"`
	}
}

// Created
// swagger:response created
type swaggCreatedResp struct {
	// in:body
	Body struct {
		// HTTP status code 201 - Status Created
		// example: 201
		Status int `json:"status"`
		// Request fulfilled - success is true
		// example: true
		Success bool `json:"success"`
		// Detailed success message
		// example: New resources created
		Message string `json:"message"`
	}
}

// No Content
// swagger:response noContent
type swaggNoContent struct {
	// in:body
	Body struct {
		// HTTP status code 204 - Status No Content
		// example: 204
		Status int `json:"status"`
		// Request fulfilled but nothing to return - success is true
		// example: true
		Success bool `json:"success"`
		// Detailed message
		// example: Operation completed successfully
		Message string `json:"message"`
	}
}

// Error: Bad Request
// swagger:response badReq
type swaggErrBadReq struct {
	// in:body
	Body struct {
		// HTTP status code 400 - Status Bad Request
		// example: 400
		Status int `json:"status"`
		// Request failed - success is False
		// example: false
		Success bool `json:"success"`
		// Detailed error message
		// example: Bad Request
		Message string `json:"message"`
	}
}

// Not found
// swagger:response notFound
type swaggNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 - Status Not Found
		// example: 404
		Status int `json:"status"`
		// No result found - success is False
		// example: false
		Success bool `json:"success"`
		// Detailed message
		// example: Nothing found
		Message string `json:"message"`
	}
}

// Error: Interval Server
// swagger:response errInternal
type swaggErrInternal struct {
	// in:body
	Body struct {
		// HTTP status code 500 - Internal server error
		// example: 500
		Status int `json:"status"`
		// Request failed - success is False
		// example: false
		Success bool `json:"success"`
		// Detailed error message
		// example: Internal server error
		Message string `json:"message"`
	}
}

// An EmployeeSampleReq model
//
// This is used for as a sample request body for update employee and create employee.
// swagger:parameters
type EmployeeSampleReq struct {
	// example: false
	Active bool `json:"active"`
	// example: T'Challa
	FirstName string `json:"firstname"`
	// example:N/A
	LastName string `json:"lastname"`
	// example: (555) 555-5555
	CellPhone string `json:"cellphone"`
	// example: Black Panther
	Title sql.NullString `json:"title"`
	// example: king
	Samaccountname string `json:"samaccountname"`
	// example: king@wakandamail.com
	Mail string `json:"mail"`
	// example: OPS
	PrimaryPa string `json:"primarypa"`
	// example: Wakanda
	Office string `json:"office"`
	// example: tchaka
	ManagerSamaccountname string `json:"managersamaccountname"`
	// example: Potentially
	TravelPref string `json:"travelpref"`
}
