package calc

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by zero is error")
	}
	return float64(a) / float64(b), nil
}

func Multiply(a, b int) int {
	return a * b
}
