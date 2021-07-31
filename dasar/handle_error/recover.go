package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("terjadi error", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Aplikasi tetap berjalan")
}