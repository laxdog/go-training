package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func area2(l int, b int) int {
	return l * b
}

func main() {
	fmt.Println(area(20, 30))
	fmt.Println(area2(20, 30))
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)
}
