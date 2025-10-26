package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))          // Output: 6
	printArgs("Alice", 30, 90, 85, 88) // Output: Name: Alice, Age: 30, Scores: [90 85 88]
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func printArgs(str string, age int, scores ...int) {
	fmt.Println("Name:", str)
	fmt.Println("Age:", age)
	fmt.Println("Scores:", scores)
}