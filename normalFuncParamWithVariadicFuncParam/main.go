package main

import "fmt"

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovered from panic!")
		}
	}()

	fmt.Println("Rectangle", 20, 30)
	fmt.Println("Square", 20)
}

func calculation(str string, y ...int) int {
	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}
