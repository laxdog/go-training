package main

import (
	"fmt"
)

func main() {
	fruits := []string{"banana", "watermelon", "apple", "coconut"}

	for i, item := range fruits {
		fmt.Printf("Fruit index %d is %s\n", i, item)
	}
}
