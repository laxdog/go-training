package main

import (
	"fmt"
	"reflect"
)

var (
	product = "Mobile"
	quantity = 50
	price = 50.50
	inStock = true
)

func main() {
	fmt.Println(product)
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(inStock)
	fmt.Println(reflect.TypeOf(product))
	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(price))
	fmt.Println(reflect.TypeOf(inStock))
}
