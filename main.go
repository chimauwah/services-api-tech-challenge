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
	_ "github.com/mattn/go-sqlite3"
)

const (
	// name of the db driver
	sqlite3Str = "sqlite3"
	// name of datasource that gets created as a successful result of open operation
	memStr = ":memory:"
)

func main() {
	// start by opening db
	db, err := InitDB()

	// load schema from file
	dot, err := LoadSQLFile("schema.sql")

	// run schema scripts
	ExecuteScript(db, dot, "create-employee-table")
	ExecuteScript(db, dot, "create-practicearea-table")
	ExecuteScript(db, dot, "create-checkin-table")
	ExecuteScript(db, dot, "create-coreskill-table")
	ExecuteScript(db, dot, "create-device-table")

	// load seed date from file
	dot, err = LoadSQLFile("seed.sql")

	// run seed data scripts
	ExecuteScript(db, dot, "insert-employee")
	ExecuteScript(db, dot, "insert-practicearea")
	ExecuteScript(db, dot, "insert-checkin")
	ExecuteScript(db, dot, "insert-coreskill")
	ExecuteScript(db, dot, "insert-coreskill2")

	//test reading an item
	stmt := "SELECT id, first_name, last_name FROM Employee WHERE office = ? and active = ?"
	rows, err := db.Query(stmt, "Charlotte", 0)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var (
		id    int
		fname string
		lname string
	)

	cnt := 0
	for rows.Next() {
		err := rows.Scan(&id, &fname, &lname)
		if err != nil {
			fmt.Printf("error occurred running (%s): (%s)\n", stmt, err)
		}
		fmt.Printf("ID: (%d) | NAME: (%s %s) \n", id, fname, lname)
		cnt++
	}
	fmt.Printf("Total rows returned: (%d)\n", cnt)

	// stmt, err := dot.Prepare(db, "drop-users-table")
	// restul, err := stmt.Exec()

	//init mux router
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
