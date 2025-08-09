package main

import (
	"fmt"
	"learn/calc"
)

func main() {
	var a int
	var b int
	var operation string

	fmt.Println("Input first number: ")
	fmt.Scan(&a)
	fmt.Println("Input second number: ")
	fmt.Scan(&b)
	fmt.Println("Input character: ")
	fmt.Scan(&operation)

	switch operation {
	case "/":
		result, err := calc.Divide(a, b)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			break
		}
		fmt.Printf("Result: %v\n", result)
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
