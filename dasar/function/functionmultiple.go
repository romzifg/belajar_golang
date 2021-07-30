package main

import "fmt"

func getFullName() (firstName string, lastName string, age int) {
	firstName = "Romzi"
	lastName = "Ghozi"
	age = 23

	return
}

func main() {
	firstName, lastName, age := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}