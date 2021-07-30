package main

import "fmt"

func main() {

	// pembuatan array cara pertama
	var names [3]string

	names[0] = "Romzi"
	names[1] = "Farhan"
	names[2] = "Ghozi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// pembuatan array cara kedua

	var angka = [4]int{
		2,
		3,
		7,
		12,
	}

	fmt.Println(angka)
	fmt.Println(len(angka))
}