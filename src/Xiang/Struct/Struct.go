package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x, y float64
}

type People struct {
	name string
}

type Employee struct {
	People     // Is-a relationship. Employee is a People
	ComanyName string
}

func main() {
	tempCir := Circle{0, 0, 5}
	fmt.Printf("The area of the circle is %f\n", tempCir.Area())
	tempRec := Rectangle{3, 4}
	fmt.Printf("The area of the rectangle is %f\n", tempRec.Area())
	fmt.Printf("The total area of the circle and rectangle is %f\n", TotalArea(tempCir, tempRec)) // Circle&Rectangle have inherited Shape interface implicitly.
	Ross := Employee{People{name: "Ross"}, "FairHR"}
	// Ross := new(Employee)
	// Ross.People.name = "Ross"
	Ross.People.SayMyName()
}

func (cir Circle) Area() float64 { // Parameter will copy variable address and copy.
	return math.Pi * cir.r * cir.r
}

func (rec Rectangle) Area() float64 {
	return rec.x * rec.y
}

func (pe *People) SayMyName() {
	fmt.Printf("Hi, my name is %s\n", pe.name)
}
