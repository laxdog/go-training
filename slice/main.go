package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randInt := make([]int, 3)

	rand.Seed(123)
	// insert some random data

	for i := 0; i < 3; i++ {
		randInt[i] = rand.Intn(100)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("%d\t", randInt[i])
	}
}
