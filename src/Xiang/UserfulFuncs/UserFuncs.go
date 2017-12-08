package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	var numList list.List
	numList.PushBack(3)
	numList.PushBack(2)
	numList.PushBack(1)
	numArray := make(IntList, numList.Len())
	arrayIndex := 0

	for item := numList.Front(); item != nil; item = item.Next() {
		if integ, ok := item.Value.(int); ok {
			numArray[arrayIndex] = integ
			arrayIndex++
		}
	}
	// sort.Sort(listItem)
	fmt.Printf("The array is %v\n", numArray)
	sort.Sort(numArray)
	fmt.Printf("The ascending sorted int array is %v\n", numArray)
}

type IntList []int

func (this IntList) Len() int {
	return len(this)
}

func (this IntList) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
