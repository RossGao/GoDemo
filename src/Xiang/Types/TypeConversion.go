package main

import (
	"fmt"
)

func ValidateType(targetItem interface{}) {
	if converted, ok := targetItem.(string); ok {
		fmt.Printf("The value is of type %T\n", converted)
		fmt.Scan()
	} else {
		fmt.Printf("The value is not type of %T\n", "")
	}
}
