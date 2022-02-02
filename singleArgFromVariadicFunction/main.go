package main

import "fmt"

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovered from panic!")
		}
	}()

	variadicExample("Red", "Blue", "Green", "Yellow")
	variadicExample("Red", "Blue", "Green", "Yellow") // Doesn't run
}

func variadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
	fmt.Println(s[4])
	fmt.Println(s[2]) // Doesn't run
}
