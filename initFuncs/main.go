package main

import "fmt"

var name string

func init() {
	fmt.Println("This is called on main init")
	name = "Dog"
}

func main() {
	fmt.Println("Start of actual main func")
	fmt.Printf("Name: %s\n", name)
}
