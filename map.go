package main

import "fmt"

func main() {
	var months map[int]string = map[int]string {
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	fmt.Println(months)
	fmt.Println(months[1])

	// Length of the map
	fmt.Println("Length of months map:", len(months))

	// Gantiing nilai pada map
	months[1] = "Jan"
	fmt.Println("Updated month for 1:", months[1])

	// Menambahkan elemen baru ke dalam map
	months[13] = "Undecimber"
	fmt.Println("Added month 13:", months[13])
	fmt.Println(months)

	// Delete elemen dari map
	delete(months, 2)
	fmt.Println("After deleting month 2:", months)
}