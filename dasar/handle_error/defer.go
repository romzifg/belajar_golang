package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runAplication(value int) {
	defer logging()
	fmt.Println("Run Aplication")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runAplication(5)
}