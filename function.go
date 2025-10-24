package main

import "fmt"

func main() {
	fmt.Println(tambah(10, 20))
	fmt.Println(kurang(10, 20))

	void(10, 20)

	angka, nama := multipleReturn()
	fmt.Println(angka)
	fmt.Println(nama)

	// Ignore return value
	angka2, _ := multipleReturn()
	fmt.Println(angka2)
}

func tambah(a int, b int) int {
	return a + b
}

func kurang(a int, b int) int {
	return a - b
}

func void(a int, b int) {
	fmt.Println(a + b)
}

func multipleReturn() (int, string) {
	return 10, "Hello"
}