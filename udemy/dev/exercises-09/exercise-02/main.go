package main

import "fmt"

type Person struct {
	name, lastName string
	age            int
}

func (p *Person) talk() {
	fmt.Println("Hi, I'm ", p.name, p.lastName, "and I'm", p.age, "years old")
}

type Human interface {
	talk()
}

func main() {
	fmt.Println("exercise 2")

	p := Person{
		name:     "Benjamin",
		lastName: "Gatica",
		age:      26,
	}

	fmt.Println(p)

	saySomething(&p)
}

func saySomething(h Human) {
	h.talk()
}
