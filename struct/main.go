package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName string
		age int
	}

	var x Employee
	x.firstName = "Tom"
	x.lastName = "Williams"
	x.age = 50

	fmt.Println(x)
}
