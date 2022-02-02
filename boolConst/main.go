package main

import "fmt"

func main() {
	const trueConst = true

	// Type definition using type keyword
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst

	// defaultBool = customBool 
	fmt.Println(defaultBool)
	fmt.Println(customBool)
}
