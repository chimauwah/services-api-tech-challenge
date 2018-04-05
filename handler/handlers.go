package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"services-api-tech-challenge/db"
	"services-api-tech-challenge/model"
	"strconv"

	"github.com/gorilla/mux"
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
	emps := db.FindAllEmployees()
	json.NewEncoder(w).Encode(emps)
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
	log.Println("CreateEmployee called...\n")
	decoder := json.NewDecoder(r.Body)
	var e model.Employee
	err := decoder.Decode(&e)
	if err != nil {
		log.Fatal(err)
		response := createResponseMap(true, "400", err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}

	emp := db.AddEmployee(e)
	json.NewEncoder(w).Encode(emp)
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
	var response map[string]string
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	cnt, err := db.DeleteEmployee(id)
	if err != nil {
		response = createResponseMap(false, "500", err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	if cnt < 1 {
		response = createResponseMap(false, "200", "Could not delete record")
	} else {
		response = createResponseMap(true, "200", "Record deleted successfully.")
	}
	json.NewEncoder(w).Encode(response)

}

func createResponseMap(success bool, code string, msg string) map[string]string {
	res := map[string]string{
		"Success": strconv.FormatBool(success),
		"Status":  code,
		"Message": msg}
	return res
}
