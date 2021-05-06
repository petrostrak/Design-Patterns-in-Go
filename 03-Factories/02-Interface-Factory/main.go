package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	Name string
	Age  int
}

type tiredPerson struct {
	Name string
	Age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I a, %d years old.\n", p.Name, p.Age)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Hi, I'm too tired to talk to you.")
}

// Because we returning an interface type (Person), we don't need
// to include the poiner *, but we need to return an & in the function.
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	// use a constructor
	p2 := NewPerson("Peter", 35)
	// p2.Age ++ // not accessible! A neat way of encapsulating info and having the factory expose only the interface
	p2.SayHello()
}
