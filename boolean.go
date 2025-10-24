package main

import "fmt"

func main() {
	fmt.Println("benar", true)
	fmt.Println("salah", false)

	var a bool = true
	var b bool = false

	fmt.Println(a || b) // true

	var c string
	if a {
		c = "hello"
	} else {
		c = "world"
	}
	fmt.Println(c)
}