package main

import "fmt"
import "net/http"
import "encoding/json"
import "bytes"

type Employee struct {
	Id   int
	Name string
}

type EmployeeDetail struct {
	Id, Age                       int
	Name, Depart, Position, Email string
}

func main() {
	emp := Employee{Id: 308869, Name: "Ross"}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(emp)
	fmt.Println("Press any key to send request...")
	var sendSignal string
	_, err := fmt.Scanf("%f\n", &sendSignal)
	res, err := http.Post("http://localhost:8080/employee", "Content-Type:application/json", buf)
	if err != nil {
		fmt.Println(err.Error())
	}
	var empDetail EmployeeDetail
	json.NewDecoder(res.Body).Decode(&empDetail)
	fmt.Printf("The employee %s's detail is:\n", emp.Name)
	fmt.Println(empDetail)
	// io.Copy(os.Stdout, res.Body)
}
