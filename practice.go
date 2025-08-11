package main

import "fmt"

func main() {

}

// Вывод выадратов от 1 до 10 включительно
func sqrt() {
	for i := 1; i < 11; i++ {
		fmt.Println(i * i)
	}
}

// Вывод суммы чисел от а до б
func rangeSum(a, b int) int {
	var sum int = 0
	for i := a; i < b+1; i++ {
		sum += i
	}
	return sum
}
