package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value (value di copy ke variable lain)
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Subang"

	fmt.Println(address1)
	fmt.Println(address2)

	// pass by reference dengan pointer (dengan menambahkan simbol &)
	address3 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address4 := &address3

	address4.City = "Bekasi"

	*address4 = Address{"Malang", "Jawa Timur", "Indonesia"}
	
	fmt.Println(address3)
	fmt.Println(address4)

	var address5 *Address = new(Address)
	address5.City = "Jakarta"
	fmt.Println(address5)

}