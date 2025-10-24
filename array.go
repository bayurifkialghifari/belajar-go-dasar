package main

func main() {
	var array [3]string = [3]string{"Hello", "World", "!"}

	var array2 [3]string
	array2[0] = "Hello"
	array2[1] = "World"
	array2[2] = "!"

	println(array[0], array[1], array[2])

	println(array2[0], array2[1], array2[2])

	// Length of array
	println("Length of array:", len(array))
}