package main

func logging() {
	println("Selesai memanggil function")
}

func main() {
	defer logging()
	println("Starting function")
}