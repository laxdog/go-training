package mylib

import "fmt"

func init() {
	fmt.Println("Init from mylib")
}

func Test() string {
	fmt.Println("Test from mylib")
	return "Hello"
}

func Check() {
	fmt.Println("Another func")
}
