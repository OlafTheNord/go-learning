package main

import (
	"fmt"
	"learn/calc"
)

func main() {
	var a int
	var b int
	var char string

	fmt.Println("Input first number: ")
	fmt.Scan(&a)
	fmt.Println("Input second number: ")
	fmt.Scan(&b)
	fmt.Println("Input character: ")
	fmt.Scan(&char)

	switch char {
	case "/":
		fmt.Printf("Result: %v\n", calc.Divine(a, b))
	case "*":
		fmt.Printf("Result: %v\n", calc.Multiply(a, b))
	case "+":
		fmt.Printf("Result: %v\n", calc.Add(a, b))
	case "-":
		fmt.Printf("Result: %v\n", calc.Subtract(a, b))
	default:
		fmt.Println("Character is wrong")
	}

}
