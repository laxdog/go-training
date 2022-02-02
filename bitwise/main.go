package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 3

	result := 0

	result = (x & y)
	fmt.Println(x, "&", y, "=", result)

	result = (x | y)
	fmt.Println(x, "|", y, "=", result)

	result = (x ^ y)
	fmt.Println(x, "^", y, "=", result)

	result = (x << 2)
	fmt.Println(x, "<<", 2, "=", result)

	result = (x >> 2)
	fmt.Println(x, ">>", 2, "=", result)
}
