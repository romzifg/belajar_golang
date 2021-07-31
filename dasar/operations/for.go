package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for dengan statement
	for num := 1; num <= 15; num++ {
		fmt.Println(num)
	}

	// for range
	slice :=[]string{"Romzi", "Farhan", "Ghozi", "Kaka"}

	// for range cara pertama
	for i:=0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range cara kedua
	for _, value := range slice {
		// fmt.Println("Index", i, "= ", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Romzi"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}