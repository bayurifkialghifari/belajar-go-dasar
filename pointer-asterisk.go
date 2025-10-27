package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "Alice", age: 30}
	// Kalau begini, maka p2 itu adalah salinan dari p1
	p2 := &p1

	// Gabakal ngaruh ke p1
	// p2 = Person{name: "Bob", age: 25}

	fmt.Println(p1)
	fmt.Println(p2)

	// Kalau mau p1 nya ikut berubah, kita harus pake pointer
	*p2 = Person{name: "Charlie", age: 28}

	fmt.Println(p1)
	fmt.Println(p2)
}