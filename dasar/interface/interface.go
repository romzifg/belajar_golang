package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// implementasi interface
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

// person mempunyai kontrak ke GetName()
func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var data Person
	data.Name = "Romzi"
	sayHello(data)

	binatang := Animal {
		Name : "Leo",
	}

	sayHello(binatang)
}