package main

import "fmt"

func main() {
	umur := 19

	switch umur {
	case 17:
		fmt.Println("Remaja")
	case 20:
		fmt.Println("Remaja Atas")
	case 25:
		fmt.Println("Dewasa")
	case 60:
		fmt.Println("Lansia")
	}	

	// short statement
	name := "Romzi"

	switch length := len(name); length > 5{
	case true:
		fmt.Println("Nama Kepanjangan")
	case false:
		fmt.Println("Nama sudah")
	}

	// switch tanpa kondisi
	age := umur
	switch {
	case age > 17 || age < 20:
		fmt.Println("Remaja")
	case age > 20 || age < 30:
		fmt.Println("Dewasa")
	case age > 30 || age < 45:
		fmt.Println("Tua")
	case age > 50:
		fmt.Println("Lansia")
	default:
		fmt.Println("Data Salah")
	}
}