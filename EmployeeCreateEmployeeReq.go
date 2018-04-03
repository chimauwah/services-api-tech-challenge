package main

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
