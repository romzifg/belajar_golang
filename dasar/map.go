package main

import "fmt"

func main() {

	// cara kesatu
	person := map[string]string{
		"name":    "Romzi Farhan",
		"address": "Bandung",
	}

	person["title"] = "Programmer"

	fmt.Println(person, len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// cara kedua
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Fundamental Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["ups"] = "Salah"
	fmt.Println(book, len(book))

	delete(book, "ups")
	fmt.Println(book, len(book))
}