package main

import "fmt"

func main() {
	p(hallo, "Asep")
}

func p(say func (string) string, name string) {
	fmt.Println(say("Hello " + name))
}

func hallo(p string) string {
	return p
}