package main

import "fmt"

func main() {
	testNum1 := 5
	SetValue(&testNum1)
	fmt.Println(testNum1)
	testNum2 := new(int)
	*testNum2 = 10
	SetValue(testNum2)
	fmt.Println(*testNum2)

	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}

func SetValue(numLocation *int) {
	*numLocation = 0
}
