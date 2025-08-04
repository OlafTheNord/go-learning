package main

import "fmt"

func main() {
	var age int = 17
	if age <= 17 {
		fmt.Printf("Your age = %d. It is smaller then we can entry to our work\n", age)
	} else {
		fmt.Println("Your age is OK for working")
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Next i = %d\n", i)
	}

	var score int = 60

	if score < 100 && score >= 90 {
		fmt.Println("S grade")
	} else if score < 90 && score >= 50 {
		fmt.Println("B grade")
	} else {
		fmt.Println("F grade")
	}

	var n int = 10
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Sum for count = %d is %d\n", n, sum)

}
