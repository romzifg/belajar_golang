package main

import "fmt"

func random() interface{} {
	return 10
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string: 
		fmt.Println("Value", value)
	case int:
		fmt.Println("Value", value)
	default:
		fmt.Println("Unkwon Type")
	}
}