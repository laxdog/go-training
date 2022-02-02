package main

import (
	"fmt"
	"reflect"
)

var s4 = "Japan"

func main() {
	var i int = 10
	var s string = "Canada"
	i2 := 3
	s2 := "Dog"
	var i3 = 4
	var s3 = "Cat"

	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(i2)
	fmt.Println(s2)
	fmt.Println(reflect.TypeOf(i3))
	fmt.Println(reflect.TypeOf(s3))

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)

	fmt.Println(s4)
	x := true

	if x {
		y := 1
		if x != false {
			fmt.Println(s4)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}
