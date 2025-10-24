package main

import "fmt"

func main() {
	var name1 string = "GGWP"
	var name2 string = "Hello"

	fmt.Println(name1)
	fmt.Println(name2)

	fmt.Println("Panjang dari name1:", len(name1)) // len() buat ngitung panjang string
	fmt.Println("Panjang dari name2:", len(name2))

	fmt.Println(name1[0]) // Dapet byte dari char index ke 0
	fmt.Println(name2[1]) // Dapet byte dari char index ke 1
}