package main

import "fmt"

func main() {
	const A = "LGH"
	var B = "LearnGoHere"

	// Concat strings
	var C = A + " " + B
	C += `!`
	fmt.Println(C)

	// Compare strings
	fmt.Println(A == "LGH")
	fmt.Println(B < A)
}
