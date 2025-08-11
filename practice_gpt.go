package main

import (
	"fmt"
	"learn/lessons"
)

func main() {
	var (
		a int
		b int
	)
	fmt.Println("Input first num")
	fmt.Scan(&a)
	fmt.Println("Input second num")
	fmt.Scan(&b)

	fmt.Printf("First number is even: %v\n", lessons.IsEven(a))
	fmt.Printf("Second number is even: %v\n", lessons.IsEven(b))
	if a == b {
		fmt.Printf("Your numbers are identical. First number %v equal second number %v", a, b)
	} else {
		fmt.Printf("Max number between first and second numers is: %v", lessons.Max(a, b))
	}
}
