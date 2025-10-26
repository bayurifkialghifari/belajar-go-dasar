package main

func random() any {
	return 123
}

func main() {
	var data any = random()

	// Type assertion to convert 'data' to int
	number, ok := data.(int)
	if ok {
		println("Data is an integer:", number)
	} else {
		println("Data is not an integer")
	}

	// Type assertion to convert 'data' to string
	text, ok := data.(string)
	if ok {
		println("Data is a string:", text)
	} else {
		println("Data is not a string")
	}
}