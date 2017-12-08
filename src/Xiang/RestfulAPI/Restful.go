package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type People struct {
	Name, Email string // The capital named field is conpulsary for Json encoding and decoding.
	Age, Phone  int
}

type Department struct {
	Name, Location string
	Id, Count      int
}

type employee struct {
	People
	Department
	Position string
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The request's path is %q", r.URL.EscapedPath())
}

func GetEmployeeId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "The employee's id is %q", vars["id"])
}

func GetDepartmentName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "The employee's department is %s", vars["deName"])
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	ross := employee{
		People{Name: "Ross", Email: "ross.gao@fairhr.com", Age: 32, Phone: 1802},
		Department{Name: "Dev Center", Location: "Futian, Shenzhen", Id: 11, Count: 12},
		"Dev"}
	err := json.NewEncoder(w).Encode(ross)
	// js, err := json.Marshal(ross)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	employeeRouter := myRouter.PathPrefix("/employee").Subrouter()
	employeeRouter.HandleFunc("/", RequestHandler)
	employeeRouter.HandleFunc("/personal/{id:[0-9]+}", GetEmployeeId)
	employeeRouter.HandleFunc("/department/{deName:[a-zA-Z]+}", GetDepartmentName)
	employeeRouter.HandleFunc("/detail", GetEmployee)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
