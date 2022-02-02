package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100000000000000"
	// strVar := "100"
	// strVar := 100
	// strVar := "100a"

	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar1, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1))

	intVar1, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1))

	intVar1, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1))

	intVar1, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1))

	intVar1, err = strconv.ParseInt(strVar, 8, 64)
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1))
}
