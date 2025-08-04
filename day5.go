package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s. I'm a %d years old.\n", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	p := Person{Name: "John", Age: 18}
	q := Person{Name: "Jane", Age: 15}

	fmt.Println(p.Greet())
	fmt.Println(p.IsAdult())
	fmt.Println(q.Greet())
	fmt.Println(q.IsAdult())
}
