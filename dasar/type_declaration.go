package main

import "fmt"

func main() {

	type NoHp string
	type Married bool

	var noRomzi NoHp = "081234567890"
	var status Married = true
	fmt.Println(noRomzi)
	fmt.Println(status)
}