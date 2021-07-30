package main

import "fmt"

func main() {
	name := "asdfaghqeoandflas"

	if name == "Romzi" {
		fmt.Println("nama depan")
	} else if name == "Farhan" {
		fmt.Println("nama tengah")
	} else if name == "Ghozi" {
		fmt.Println("nama belakang")
	} else {
		fmt.Println("data salah")
	}

	// short statement if
	if length := len(name); length > 10 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah Benar")
	}
}