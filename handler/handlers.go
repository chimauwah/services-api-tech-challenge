package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"services-api-tech-challenge/model"
)

// GetEmployees swagger:route GET /api/employees employees listEmployees
//
// Resource returning all employees.
//
// Responses:
//	200: employeeResponse
//	500: internal
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	log.Print("GetEmployees called...\n")
	emp := new(model.Employee)
	json.NewEncoder(w).Encode(emp)
}

// GetEmployeeDetails swagger:route GET /api/employee/details/{id} employeedetails listEmployeeDetails
//
// Resource returning details for a specific employee.
//
// Responses:
//	200: employeeResponse
//	500: internal
func GetEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	log.Print("GetEmployeeDetails called...")
}

// CreateEmployee swagger:route POST /api/employee employee createEmployee
//
// Resource to create a single employee.
//
// Responses:
//	200: employeeResponse
//  400: badReq
//  422: validationError
//	500: internal
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	log.Print("CreateEmployee called...")
}

// UpdateEmployee swagger:route PUT /api/employee/{id} employee updateEmployee
//
// Resource to update an existing employee.
//
// Responses:
//	200: employeeResponse
//  400: badReq
//  422: validationError
//	500: internal
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	log.Print("UpdateEmployee called...\n")
}

// DeleteEmployee swagger:route DELETE /api/employee/{id} employee deleteEmployee
//
// Resource to delete an existing employee.
//
// Responses:
//	200: employeeResponse
//  400: badReq
//	500: internal
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteEmployee called...")
}
