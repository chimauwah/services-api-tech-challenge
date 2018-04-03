package main

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
