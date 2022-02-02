package main

import "fmt"

func main() {
	A := [3][3]int{
		{1, 1, 0},
		{1, 1, 1},
		{1, 0, 0},
	}
	B := [3][3]int{
		{0, 1, 0},
		{1, 0, 1},
		{1, 0, 1},
	}
	var C [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				C[i][j] = C[i][j] + A[i][k] * B[k][i]
			}
		}
	}
	fmt.Println("A = ", A)
	fmt.Println("B = ", B)
	fmt.Println("C = ", C)
}
