package main

import "fmt"

func main() {
	var integer32 int32 = 32345234
	var integer64 int64 = int64(integer32)

	// Kalau angka lebih dari kapasitas int16, maka akan terjadi overflow
	var integer16 uint16 = uint16(integer32)


	fmt.Println("integer32:", integer32)
	fmt.Println("integer64:", integer64)
	fmt.Println("integer16:", integer16)


	var asep string = "Asep"
	var aint uint8 = asep[0]
	var a string = string(aint)

	fmt.Println("asep:", asep)
	fmt.Println("aint:", aint)
	fmt.Println("a:", a)
}