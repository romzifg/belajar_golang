package main

import "fmt"

func main() {
	var a = 12
	var b = 6
	var c = 9

	var result1 = a + b
	var result2 = a - c
	var result3 = a * b
	var result4 = a / b
	var result5 = a % c

	// augmented
	a += c
	a++

	fmt.Println(result1, result2, result3, result4, result5)
	fmt.Println(a)
}