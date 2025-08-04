package main

import "fmt"

func main() {
	var name string = "Phil"
	var age int = 37
	var height float64 = 1.95

	fmt.Printf("My name %s, im %d age, my hrght %.2f m.\n", name, age, height)

	const PI float64 = 3.1415
	var radius float64 = 4.0
	var area = PI * radius * radius

	fmt.Printf("S of with R = %.1f: %.2f\n", radius, area)
}
