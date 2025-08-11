package slicer

import (
	"fmt"
)

func Slicer(a int) (int, int) {
	var (
		slice    []int
		multiply int
		sum      int
		number   int
	)
	for i := a; i > 0; i-- {
		fmt.Println("Input a number: ")
		fmt.Scan(&number)
		slice = append(slice, number)
	}
	multiply = 1
	for _, c := range slice {
		sum += c
		multiply *= c
	}

	return sum, multiply
}
