package main

import "fmt"

func main() {
	empSalary := make(map[string]int)
	empSalary["steve"] = 12000
	empSalary["james"] = 15000
	empSalary["mike"] = 9000

	fmt.Println("map contents: ", empSalary)

	empSalary2 := map[string]int{
		"steve" : 12345,
		"james" : 23456,
		"mike" : 34567,
	}
	emp := "mike"
	sal := empSalary2[emp]
	fmt.Println("Salary of ", emp, "is", sal)
	delete(empSalary2, emp)
	fmt.Println(empSalary2)
	fmt.Println(empSalary2[emp])
}

