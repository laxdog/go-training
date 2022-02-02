package main

import (
	"fmt"
)

type tank interface {
	// Methods
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

// Implement methods of tank interface
func (m myvalue) Tarea() float64 {
	return 2 * m.radius * m.height * +2 * 3.14 * m.radius * m.radius
}

func (m myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

// Main
func main() {
	// Accessing elements of the tank interface
	var t tank
	fmt.Println("Value :", t)
	fmt.Println("Type %T", t)
	t = myvalue{10, 14}

	fmt.Println("Value :", t)
	fmt.Println("Type %T", t)
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank ", t.Volume())
}
