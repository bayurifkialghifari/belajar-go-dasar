package main

import "fmt"

func main() {
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 []string = months[4:7] // ambil nya dari index 4 sampai 6
	fmt.Println("slice1:", slice1)

	slice2 := months[10:]
	fmt.Println("slice2:", slice2)

	slice3 := months[:3]
	fmt.Println("slice3:", slice3)

	slice4 := months[:]
	fmt.Println("slice4:", slice4)

	// Append slice
	slice5 := []string{"GG", "WP", "!"}
	fmt.Println("slice5 sebelum append:", slice5)

	slice5 = append(slice5, "EZ")
	fmt.Println("slice5 setelah append:", slice5)

	// Remove slice
	slice6 := []string{"G", "G", "W", "P", "!"}
	fmt.Println("slice6 sebelum di remove:", slice6)

	fmt.Println(slice6[:2])
	// Jadi index ke 2 itu di skip
	fmt.Println(slice6[3:])

	slice6 = append(slice6[:2], slice6[3:]...) // menghapus index ke 2
	fmt.Println("slice6 setelah di remove:", slice6)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "GG"
	newSlice[1] = "WP"
	// newSlice[100] = "!" // Akan menyebabkan panic: runtime error: index out of range [100] with length 5

	fmt.Println("newSlice:", newSlice)
	fmt.Println("len newSlice:", len(newSlice))
	fmt.Println("cap newSlice:", cap(newSlice))
}