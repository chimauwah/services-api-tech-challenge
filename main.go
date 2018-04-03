// Services & Api Tech Challange #1
//
// This Rest API written in GoLang provides CRUD operations
// for a directory of company employees
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// init mux router
	router := mux.NewRouter()

	// route handlers / endpoints
	router.HandleFunc("/api/employees", getEmployees).Methods("GET")
	router.HandleFunc("/api/employee/details/{id}", getEmployeeDetails).Methods("GET")
	router.HandleFunc("/api/employee", createEmployee).Methods("POST")
	router.HandleFunc("/api/employee/{id}", updateEmployee).Methods("PUT")
	router.HandleFunc("/api/employee/{id}", deleteEmployee).Methods("DELETE")

	fmt.Println("Listening on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}

// GetEmployees swagger:route GET /api/employees employees listEmployees
//
// Resource returning all employees.
//
// Responses:
//	200: employeeResponse
//	500: internal
func getEmployees(w http.ResponseWriter, r *http.Request) {

}

// GetEmployeeDetails swagger:route GET /api/employee/details/{id} employeedetails listEmployeeDetails
//
// Resource returning details for a specific employee.
//
// Responses:
//	200: employeeResponse
//	500: internal
func getEmployeeDetails(w http.ResponseWriter, r *http.Request) {

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
func createEmployee(w http.ResponseWriter, r *http.Request) {

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
func updateEmployee(w http.ResponseWriter, r *http.Request) {

}

// DeleteEmployee swagger:route DELETE /api/employee/{id} employee deleteEmployee
//
// Resource to delete an existing employee.
//
// Responses:
//	200: employeeResponse
//  400: badReq
//	500: internal
func deleteEmployee(w http.ResponseWriter, r *http.Request) {

}
