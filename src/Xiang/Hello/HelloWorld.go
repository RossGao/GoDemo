package main

import (
	"StringUtil"
	"fmt"
)

func main() {
	fmt.Printf(stringUtil.Reverse("\n.dlrow ,olleH"))
	defer func() {
		err := recover() // defer is a function call. Always been calld before function return
		fmt.Println(err)
	}()
	total := Sum(1, 2, 3, 4, 5)
	if total == 0 {
		panic("Test panic message")
	}
	fmt.Printf("Total number is %d\n", total)
	half, isEven := HalfAndEven(3)
	fmt.Printf("%d's half is %d which is %v\n", 3, half, isEven)
	// err := recover()		Never run
	// fmt.Println(err)
	testNum := 109
	firstNum, secondNum := 0, 1
	isFab, Seq := FabonacciSequence(firstNum, secondNum, testNum)
	if isFab {
		resultSeq := append([]int{firstNum}, Seq...)
		fmt.Printf("%d is Fabonacci number, the sequence is %d\n", testNum, resultSeq)
	} else {
		fmt.Printf("%d is not a Fabonacci number\n", testNum)
	}
}

// Sum up function for varia number of integers.
func Sum(items ...int) int {
	total := 0
	for _, i := range items {
		total += i
	}

	return total
}

// Return two values
func HalfAndEven(ori int) (int, bool) {
	if ori%2 == 0 {
		return 0, true
	}
	return 1, false
}

func FabonacciSequence(first, second, targetNum int) (bool, []int) {
	if first < targetNum && second < targetNum {
		addition := first + second
		seqPart := []int{second}

		if addition < targetNum {
			isFab, seq := FabonacciSequence(second, addition, targetNum)
			return isFab, append(seqPart, seq...)
		} else if addition == targetNum {
			return true, append(seqPart, targetNum)
		} else {
			return false, nil
		}
	}

	return false, nil
}
