package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chimauwah/services-api-tech-challenge/db"
	"github.com/chimauwah/services-api-tech-challenge/model"
	"github.com/gorilla/mux"
)

// GetEmployee swagger:route GET /api/employee/{id} employee getEmployee
//
// Retrieve employee with given id.
//
// Responses:
//	200: employeeResponse
//	500: errInternal
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer printExecTime(start)
	log.Println("GetEmployee called...")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	emp, err := db.FindEmployee(id)
	if err != nil {
		log.Println(err)
		response := createResponseMap(false, "500", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(emp)
}

// GetEmployees swagger:route GET /api/employees employees listEmployees
//
// Retrieve all employees.
//
// Responses:
//	200: employeesResponse
//	500: errInternal
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer printExecTime(start)
	log.Println("GetEmployees called...")
	w.Header().Set("Content-Type", "application/json")
	emps, err := db.FindAllEmployees()
	if err != nil {
		log.Println(err)
		response := createResponseMap(false, "500", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(emps)
}

// GetEmployeeDetails swagger:route GET /api/employee/details/{id} employeedetails listEmployeeDetails
//
// Retreive details for a specific employee.
//
// Responses:
//	200: employeeDtlResponse
//	500: errInternal
func GetEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer printExecTime(start)
	log.Println("GetEmployeeDetails called...")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	det, err := db.FindEmployeeDetails(id)
	if err != nil {
		log.Println(err)
		response := createResponseMap(false, "500", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(det)

}

// SearchEmployees swagger:route GET /api/employee/search employees searchEmployees
//
// Search for an employee with provided criteria.
//
// Responses:
//	200: employeesResponse
//  404: notFound
//	500: errInternal
func SearchEmployees(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer printExecTime(start)
	log.Println("SearchEmployees called...")
	w.Header().Set("Content-Type", "application/json")
	vals := r.URL.Query()
	emps, err := db.SearchEmployees(vals)
	if err != nil {
		log.Println(err)
		response := createResponseMap(false, "500", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	if emps == nil {
		response := createResponseMap(false, "404", "No results found")
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(emps)
}

// CreateEmployee swagger:route POST /api/employee employee createEmployee
//
// Create a single employee.
//
// Responses:
//	201: employeeResponse
//  400: badReq
//	500: errInternal
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer printExecTime(start)
	log.Println("CreateEmployee called...")
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var e model.Employee
	err := decoder.Decode(&e)
	if err != nil {
		log.Println(err)
		response := createResponseMap(false, "400", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}

	emp, err := db.AddEmployee(e)
	if err != nil {
		log.Println(err)
		response := createResponseMap(false, "500", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(emp)
}

// UpdateEmployee swagger:route PUT /api/employee/{id} employee updateEmployee
//
// Update an existing employee.
//
// Responses:
//  204: noContent
//  400: badReq
//  404: notFound
//	500: errInternal
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer printExecTime(start)
	log.Println("UpdateEmployee called...")
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var e model.Employee
	err := decoder.Decode(&e)
	if err != nil {
		log.Println(err)
		response := createResponseMap(false, "400", "Could not decode request body: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}

	var response map[string]string
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	cnt, err := db.UpdateEmployee(id, e)
	if err != nil {
		log.Println(err)
		response = createResponseMap(false, "500", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}

	if cnt < 1 {
		response = createResponseMap(false, "404", "Could not update record")
	} else {
		response = createResponseMap(true, "204", "Record updated successfully.")
	}
	json.NewEncoder(w).Encode(response)

}

// DeleteEmployee swagger:route DELETE /api/employee/{id} employee deleteEmployee
//
// Delete an existing employee.
//
// Responses:
//  204: noContent
//  404: notFound
//	500: errInternal
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer printExecTime(start)
	log.Println("DeleteEmployee called...")
	w.Header().Set("Content-Type", "application/json")
	var response map[string]string
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	cnt, err := db.DeleteEmployee(id)
	if err != nil {
		log.Println(err)
		response = createResponseMap(false, "500", "Error: "+err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	if cnt < 1 {
		response = createResponseMap(false, "404", "No record found to delete.")
	} else {
		response = createResponseMap(true, "204", "Record deleted successfully.")
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

func printExecTime(start time.Time) {
	log.Printf("Total time to execute (%s) ", time.Since(start))
}
