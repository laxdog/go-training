package main

import "fmt"

func main() {
	var intArray[5] int
	fmt.Printf("%T %d %v\n", intArray, len(intArray), intArray)


	strArray := [...]string{"Mercury", "Venus", "Earth", "Mars"}
	fmt.Printf("n%T %d %q\n", strArray, len(strArray), strArray)
}
