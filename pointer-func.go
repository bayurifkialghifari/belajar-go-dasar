package main

func pointerFunc(b *int) {
	*b += 10
}

func main() {
	var a int = 10
	println(a)
	pointerFunc(&a)
	println(a) // Output: 20
}