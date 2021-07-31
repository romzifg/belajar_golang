package main

import "fmt"

func main() {
	var nilai int32 = 130
	var nilai64 int64 = int64(nilai)
	var nilai8 int8 = int8(nilai)

	fmt.Println(nilai)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Romzi"
	var e = name[0]
	var eString = string(e)

	fmt.Println(eString)
}