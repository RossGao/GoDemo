package main

import "fmt"

func mapTested() {
	fmt.Println("Hello world!") // Just the comments
	fmt.Println("Hello world!"[1])
	fmt.Println(len("Hello world!"))
	fmt.Println("321245 * 424521 is", 321245*424521)
	fmt.Println((true && false) || (false && true) || !(false && false))
	var x string = "First"
	ValidateType(x)
	fmt.Println(x)
	y := 7
	fmt.Println(y)
	var (
		pen   = "pen"
		paper = "A4"
		ink   = "Dark blue"
		words = "Hello world"
		novel = "The dawn"
	)
	fmt.Println(pen, paper, ink, words, novel)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, "Even")
		} else {
			fmt.Println(j, "Odd")
		}
	}

	// region
	arr := [6]float64{82, 76, 78, 83, 90}
	fmt.Println(arr)

	k := []int{1, 2, 3, 4}
	l := append(k, 5)
	m := make([]int, 3)
	copy(m, k)
	fmt.Println("l", l)
	fmt.Println("m", m)
	// endregion

	cars := make(map[string]string)
	cars["sports"] = "Faraly"
	cars["Luxery"] = "Bently"
	cars["Quarlity"] = "Benze"
	cars["Suv"] = "XTrail"
	cars["Electric"] = "比亚迪"

	if car, ok := cars["Classic"]; ok {
		fmt.Println(car)
	} else {
		fmt.Println("Can't find car")
	}

	scorces := []int{92, 98, 85, 87, 86}
	fmt.Println("The average score is ", Everage(scorces))

	fmt.Println(Factorial(4))
}

func Everage(scorces []int) int {
	if len(scorces) == 0 {
		panic("Cannot be empty array")
	}

	total := 0

	for _, scorce := range scorces {
		total += scorce
	}

	return total / len(scorces)
}

func Factorial(factor int) int {
	if factor == 0 {
		return 1 // Set the end point of the factorial
	}

	return factor * (factor - 1)
}
