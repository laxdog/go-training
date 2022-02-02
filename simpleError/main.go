package main

import (
	"errors"
	"fmt"
)

func e(v int) (int, error) {
	if v == 0 {
		return 0, errors.New("Zero cannot be used")
	} else {
		return 2 * v, nil
	}
}

func main() {
	v, err := e(0)

	if err != nil {
		fmt.Println(err, v)
	}
}
