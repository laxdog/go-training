package main

import (
	"errors"
	"fmt"
)

func returnError() (int, error) {
	return 42, errors.New("Error occured!")
}

func main() {
	v, _ := returnError()

	fmt.Println(v)
}
