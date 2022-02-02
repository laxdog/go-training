package main

import "fmt"

func main () {
	var a int = 20
	var ip *int
	var ptr *int

	ip = &a

	fmt.Printf("Address of var a: %x\n", &a)

	fmt.Printf("Address stored in ip var: %x\n", ip)

	fmt.Printf("Value of *ip variable: %d\n", *ip)

	fmt.Printf("Value of ptr variable: %x\n", ptr)
	// panic, nil point dereference
	// fmt.Printf("Value of ptr variable: %d\n", *ptr)
}
