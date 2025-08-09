package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)
	height := 190.0
	var heightQuestion string
	var isMarried bool
	var marriedQuestion string
	const country = "Russia"

	fmt.Println("What is your name?")
	fmt.Scan(&name)

	fmt.Println("How old are you?")
	fmt.Scan(&age)

	fmt.Println("Your height is 190. Is it? (Y / N)")
	fmt.Scan(&heightQuestion)
	if heightQuestion == "N" {
		fmt.Println("What is your height?")
		fmt.Scan(&height)
	} else if heightQuestion == "Y" {
		fmt.Println("Ok, I knew it.")
	} else {
		fmt.Println("You press something wrong. Enter your height.")
		fmt.Scan(&height)
	}

	fmt.Println("Are you married?")
	fmt.Scan(&marriedQuestion)
	if marriedQuestion == "Y" {
		isMarried = true
	}
	for marriedQuestion != "Y" && marriedQuestion != "N" {
		fmt.Println("Please, choose between Y or N")
		if marriedQuestion == "Y" {
			isMarried = true
		}
		fmt.Scan(&marriedQuestion)
	}

	fmt.Printf("Name: %s\nAge: %v\nHeight: %v\nMarried: %v\nCountry: %s", name, age, height, isMarried, country)
}
