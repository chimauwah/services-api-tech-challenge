package main

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
