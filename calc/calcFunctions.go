package calc

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divine(a, b int) float64 {
	if b == 0 {
		fmt.Println("Divine by zero is error")
		return 0
	}
	return float64(a) / float64(b)
}

func Multiply(a, b int) int {
	return a * b
}
