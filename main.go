package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	//Go provides a flag package supporting basic command-line flag parsing.
	addr = flag.String("addr", ":8080", "http service port")
	data map[string]string
)

func main() {
	// parse command line
	flag.Parse()
	data = map[string]string{}
	router := httprouter.New()

	router.GET("/", home)
	router.GET("/employees", retrieveAllEmployees)
	router.GET("/employee/:id", retrieveEmployeeDetails)
	router.POST("/add", addEmployee)
	router.PUT("/update/:key/:value", updateEmployee)
	router.DELETE("/remove/:key", removeEmployee)

	fmt.Print("Listening on port 8080...\n")
	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func home(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Write([]byte("Welcome to Chima Uwah's Services & APIs Tech Challenge #1 completed in GoLang"))
}

func retrieveEmployeeDetails(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Write([]byte("Not yet implemented"))
}

func retrieveAllEmployees(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// k := p.ByName("key")
	// if k == "" {
	// 	// if no key in url, list all entries in data map
	// 	fmt.Fprintf(w, "Read list: %v", data)
	// 	return
	// }
	// // if key given, returns corresponding value
	// fmt.Fprintf(w, "ready entry: data[%s] = %s", k, data[k])
	w.Write([]byte("Not yet implemented"))

}

func addEmployee(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Write([]byte("Not yet implemented"))
}

func removeEmployee(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Write([]byte("Not yet implemented"))
}

func updateEmployee(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// k := p.ByName("key")
	// v := p.ByName("value")

	w.Write([]byte("Not yet implemented"))

	// data[k] = v

	// fmt.Fprintf(w, "Updated: data[%s] = %s", k, data[k])
}
