package main

import (
	"fmt"
	"learn/slicer"
)

func main() {
	var (
		counts int
		sum    int
		mult   int
	)
	fmt.Println("Enter a count of numbers")
	fmt.Scan(&counts)
	sum, mult = slicer.Slicer(counts)
	fmt.Printf("Result is\nSum: %v\nMultiply of all numbers: %v\n", sum, mult)
}
