package main

import "fmt"

func main() {
	type Umur uint8
	type KTP string

	var age Umur = 30
	var id KTP = "1234567890"

	fmt.Println("Umur saya adalah:", age)
	fmt.Println("Nomor KTP saya adalah:", id)
}