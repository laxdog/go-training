package main

import (
	"fmt"
)

func typeswitch(a interface{}) {
	// Using type switch
	switch a.(type) {
	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("Type: string, Value:", a.(string))
	case float64:
		fmt.Println("Type: float64, Value:", a.(float64))
	default:
		fmt.Printf("Type not found %+v\n", a)
	}
}

// Main
func main() {
	typeswitch("Test")
	typeswitch(67.9)
	typeswitch(true)
}
