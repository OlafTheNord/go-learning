package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
func greet(name string) string {
	return "Hello " + name
}
func main() {
	fmt.Println(multiply(1, 2))
	fmt.Println(isEven(multiply(2, 3)))
	fmt.Println(greet("John"))
}
