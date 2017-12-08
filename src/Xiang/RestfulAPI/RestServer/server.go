package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id   int
	Name string
}

type EmployeeDetail struct {
	Id, Age                       int
	Name, Depart, Position, Email string
}

func GetEmployeeDetail(w http.ResponseWriter, r *http.Request) {
	var e Employee
	var reBytes []byte
	if _, err := r.Body.Read(reBytes); err != nil { // Check if the post content is empty or not.
		http.Error(w, "Please provide post content.", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	empDetail := EmployeeDetail{Id: e.Id, Age: 32, Name: e.Name, Depart: "Dev Center", Position: "Dev", Email: "ross.gao@fairhr.com"}
	json.NewEncoder(w).Encode(empDetail)
}

func main() {
	srvRouter := mux.NewRouter().PathPrefix("/employee").Subrouter()
	srvRouter.HandleFunc("/", GetEmployeeDetail).Methods("POST") // post request. Update or create new employee.

	log.Fatal(http.ListenAndServe(":8080", srvRouter))
}
