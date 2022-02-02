package main

import "fmt"

func main() {
	fmt.Printf("[%*s]\n", 5, "abc")
	fmt.Printf("[%-*s]\n", 5, "abc")
	fmt.Printf("[%*s]\n", 3, "abcdef")
	fmt.Printf("[%*d]", 5, 123)
	fmt.Printf("[%*d]\n", 3, 12345)
	fmt.Printf("[%*.2f]\n", 5, 3.456)
	fmt.Printf("[%*.*f]\n", 5, 2, 3.456)
}
