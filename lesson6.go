package main

import (
	"fmt"
	"learn/lessons"
)

func main() {
	var (
		count   int
		sum_num int
		min_num int
		max_num int
		avg_num float64
	)
	fmt.Println("Input count of numbers: ")
	fmt.Scan(&count)
	sum_num, min_num, max_num, avg_num = lessons.ProcessStream(count)
	fmt.Printf("Result is:\nSum: %v\nMin: %v\nMax: %v\nAverage: %v\n", sum_num, min_num, max_num, avg_num)
}
