package main

import "fmt"

func main() {
	var pertambahan int = 10 + 5
	var pengurangan int = 10 - 5
	var perkalian int = 10 * 5
	var pembagian int = 10 / 5
	var modulus int = 10 % 3

	// Augmented assignment sama kaya pemgoraman lagin
	var i = 100
	i += 10  // i = i + 10
	// i -= 10  // i = i - 10
	// i *= 10  // i = i * 10
	// i /= 10  // i = i / 10
	// i %= 3   // i = i % 3

	fmt.Println("Hasil Pertambahan: ", pertambahan)
	fmt.Println("Hasil Pengurangan: ", pengurangan)
	fmt.Println("Hasil Perkalian: ", perkalian)
	fmt.Println("Hasil Pembagian: ", pembagian)
	fmt.Println("Hasil Modulus: ", modulus)
	fmt.Println("Hasil Akhir i: ", i)
}