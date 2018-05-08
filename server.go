// Services & Api Tech Challange #1
//
// This basic modern RESTful API written in GoLang, provides CRUD operations
// and search functionality for a directory of hard-coded company employees
// stored in an in-memory database
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.1.0
//	   Contact: Chima Uwah<cuwah@captechconsulting.com> https://github.com/chimauwah/services-api-tech-challenge
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
	"github.com/chimauwah/services-api-tech-challenge/middleware"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// start by initializing db
	db.Init()

	//init mux router
	router := mux.NewRouter()

	// route handlers / endpoints
	router.HandleFunc("/api/employees", handler.GetEmployees).Methods("GET")
	router.HandleFunc("/api/employee/{id:[0-9]+}", handler.GetEmployee).Methods("GET")
	router.HandleFunc("/api/employee/details/{id:[0-9]+}", handler.GetEmployeeDetails).Methods("GET")
	router.HandleFunc("/api/employee/search", handler.SearchEmployees).Methods("GET")
	router.HandleFunc("/api/employee", handler.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee/{id:[0-9]+}", handler.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/employee/{id:[0-9]+}", handler.DeleteEmployee).Methods("DELETE")

	redoc := middleware.Redoc(middleware.RedocOpts{}, router)

	router.HandleFunc("/api/docs", redoc.ServeHTTP).Methods("GET")

	fmt.Println("Listening on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
