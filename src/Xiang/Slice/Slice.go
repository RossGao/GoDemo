package main

import "fmt"

func main() {
	shortSlice := make([]int, 200)
	fmt.Printf("Size of the origin slice is %v\n", cap(shortSlice))
	newSlice := ReallocateSlice(shortSlice)
	fmt.Printf("Size of the reallocated slice is %v\n", cap(newSlice))
}

func ReallocateSlice(slice []int) (longSlice []int) {
	if cap(slice) < 101 {
		newSlice := make([]int, cap(slice)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	return slice
}
