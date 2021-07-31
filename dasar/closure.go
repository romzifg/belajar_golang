package main

import "fmt"

func main() {
	name := "Romzi"
	counter := 0

	increment := func() {
		name = "Ghozi"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}