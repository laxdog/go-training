package main

import (
	"fmt"
)

func f(s string) {
	panic(s)
}

func main() {
	// defer makes the function run at the end
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	f("Panic!")
}
