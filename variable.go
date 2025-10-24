package main

import "fmt"

func main() {
	var a = "Hello"
	b := "World"

	var c string = "Malas"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Multiple variable declaration
	var (
		firstName = "John"
		lastName  = "Doe"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}