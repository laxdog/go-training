package main

import (
	"fmt"
	"mylib/mylib"
	"mylib/mylib/alib"
)

func ToTest() string {
	return mylib.Test()
}

func main() {
	fmt.Println("This is a test")
	mylib.Test()
	mylib.Check()
	alib.ATest()
}
