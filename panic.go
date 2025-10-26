package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error:", message)
	}
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Aplikasi error")
	}
}