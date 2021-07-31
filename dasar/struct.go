package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
	Married		bool
}

// struct function/method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var data Customer

	data.Name = "Romzi"
	data.Adress = "Indonesia"
	data.Age = 23
	data.Married = false

	// pemanggilan struct function
	data.sayHello("Friska")

	// fmt.Println(data)

	// data2 := Customer {
	// 	Name : "Ghozi",
	// 	Adress : "Rancaekek",
	// 	Age : 24,
	// 	Married: false,
	// }

	// fmt.Println(data2)
}