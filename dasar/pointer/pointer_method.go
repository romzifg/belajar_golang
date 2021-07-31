package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	data := Man{"Romzi"}
	data.Married()

	fmt.Println(data.Name)
}