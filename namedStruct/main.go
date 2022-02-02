package main

import "fmt"

type Employee struct {
	firstName string
	lastName string
	age int
	salary int
}

func main() {
	// create struct specifying field names
	emp1 := Employee {
		firstName: "Same",
		age:	25,
		salary:	500,
		lastName:	"Anderson",
	}

	// create struct withoug specifying field names
	// notice the order must match
	emp2 := Employee{"Thomas", "John", 29, 800}
	fmt.Println("Emp 1", emp1)
	fmt.Println("Emp 2", emp2)
}
