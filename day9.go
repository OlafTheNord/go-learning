package main

import "fmt"

func main() {
	var firstNum int
	var secNum int
	var char string
	fmt.Println("Input first number: ")
	fmt.Scan(&firstNum)
	fmt.Println("Input second number: ")
	fmt.Scan(&secNum)
	fmt.Println("Input character: ")
	fmt.Scan(&char)

	switch {
	case char == "*":

		fmt.Printf("Result: %v", mltpl(firstNum, secNum))
	case char == "/":

		fmt.Printf("Result: %v", divine(firstNum, secNum))
	case char == "+":
		fmt.Printf("Result: %v", sum(firstNum, secNum))
	case char == "-":

		fmt.Printf("Result: %v", notSum(firstNum, secNum))
	default:

		fmt.Println("Something wrong!!!")
	}
}

func mltpl(a, b int) int {
	return a * b
}

func divine(a, b int) float64 {
	if b == 0 {
		fmt.Println("Num cant be zero")
	}
	return float64(float64(a) / float64(b))
}

func sum(a, b int) int {
	return a + b
}

func notSum(a, b int) int {
	return a - b
}
