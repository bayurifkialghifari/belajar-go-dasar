package main

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "Alice", age: 30}
	// Kalau begini, maka p2 itu adalah salinan dari p1
	p2 := p1

	p2.age = 31

	println("p1 age:", p1.age) // Output: p1 age: 30
	println("p2 age:", p2.age) // Output: p2 age: 31

	// Kalau mau p1 nya ikut berubah, kita harus pake pointer
	// p3 := &p1
	var p3 *Person = &p1 // sama aja kaya di atas

	p3.age = 32

	println("p1 age:", p1.age) // Output: p1 age: 32
	println("p3 age:", p3.age) // Output: p3 age: 32
}