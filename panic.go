package main

func main() {
	runApp(true)
}

func endApp() {
	println("Aplikasi selesai")
}

func runApp(err bool) {
	defer logging()
	if err {
		panic("Aplikasi error")
	}
}