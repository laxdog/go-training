package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func ValueReceiver(p Person) {
	p.Name = "John"
	fmt.Println("Inside value receiver", p.Name)
}

func PointerReceiver(p *Person) {
	p.Age = 25
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(*p))
	fmt.Println("Inside pointer receiver model", p.Age)
}

func main() {
	p := Person{"Tom", 28}
	p1 := &Person{"Patrick", 68}
	ValueReceiver(p)
	fmt.Println("Inside main after value receiver : ", p.Name)
	PointerReceiver(p1)
	fmt.Println("Inside main after pointer receiver : ", p1.Age)
}
