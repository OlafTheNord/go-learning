package main

import "fmt"

type Speaker interface {
	Speak() string
}
type Dog struct {
	Name string
}

type PersonNew struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Woof! My name if %s!", d.Name)
}

func (p PersonNew) Speak() string {
	return fmt.Sprintf("Hello! My name is %s!", p.Name)
}

func Introduce(s Speaker) {
	fmt.Println("I can speak " + s.Speak())
}

func main() {
	d := Dog{Name: "Bublik"}
	p := PersonNew{Name: "John"}
	Introduce(d)
	Introduce(p)
}
