package main

import "fmt"

func main() {
	var num int
	fmt.Println("Input num: ")
	fmt.Scan(&num)

	if num == 0 {
		fmt.Println("Num is 0")
	} else if num%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is not even")
	}
}
