package main

import "fmt"

func main() {

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

	slice1[0] = "Mei Lagi"
	fmt.Println(slice1)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Romzi")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)


	//  make slice
	newSlice := make([]string, 2,5)
	newSlice[0] = "Romzi"
	newSlice[1] = "Ghozi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}