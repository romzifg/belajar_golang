package main

import "fmt"

func main() {
	// deklarasi variable cara pertama
	var name string

	name = "Romzi Farhan"
	fmt.Println(name)

	name = "Ghozi"
	fmt.Println(name)

	var fullName = "Romzi Farhan Ghozi"
	fmt.Println(fullName)

	var age = 23
	fmt.Println(age)

	// deklarasi variable cara kedua
	angka := 23
	fmt.Println(angka)

	angka = 30
	fmt.Println(angka)

	// deklarasi multiple variable
	var (
		country = "Indonesia"
		province = "Jawa Barat"
	)

	fmt.Println(country)
	fmt.Println(province)
}