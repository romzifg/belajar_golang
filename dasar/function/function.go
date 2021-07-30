package main

import "fmt"

// tanpa parameter
func hitung() {
	fmt.Println("ini adalah function hitung")
}

// dengan parameter
func hitungAngka(a int, b int) {
	fmt.Println(a, "+", b, "=" , a + b)
}

// function utama
func main() {
	hitung()
	hitungAngka(2,3)
}