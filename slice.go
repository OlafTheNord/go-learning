package main

import (
	"fmt"
	"learn/slicer"
)

func main() {
	var count int
	fmt.Println("Input count of numbers: ")
	fmt.Scan(&count)
	sum, min_num, max_num, avg := slicer.Slice(count)
	fmt.Printf("Result:\nSum = %v\nMinimal number:%v\nMaximum number:%v\nAverage:%v\n", sum, min_num, max_num, avg)
}
