package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	fmt.Println(squareSum(5)(6)(7))
}
