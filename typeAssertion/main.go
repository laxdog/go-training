package main

import (
	"fmt"
)

func typeassert(a interface{}) {
	// value := a.(float64)
	value, ok := a.(float64)
	// fmt.Println(value)
	fmt.Println(value, ok)
}

// Main
func main() {

	var a1 interface{} = 98.09

	typeassert(a1)

	var a2 interface{} = "Test"

	typeassert(a2)

}
