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

	"github.com/chimauwah/services-api-tech-challenge/db"
	"github.com/chimauwah/services-api-tech-challenge/handler"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// start by initializing db
	db.Init()

	//init mux router
	router := mux.NewRouter()

	// route handlers / endpoints
	// router.Methods("GET").Path("/api/employeez").HandlerFunc(handler.GetEmployees)
	router.HandleFunc("/api/employees", handler.GetEmployees).Methods("GET")
	router.HandleFunc("/api/employee/{id}", handler.GetEmployee).Methods("GET")
	router.HandleFunc("/api/employee/details/{id}", handler.GetEmployeeDetails).Methods("GET")
	router.HandleFunc("/api/employee", handler.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee/{id}", handler.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/employee/{id}", handler.DeleteEmployee).Methods("DELETE")

	fmt.Println("Listening on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
