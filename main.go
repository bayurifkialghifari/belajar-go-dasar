package main

import (
	"fmt"
	"golang-package/helper"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(helper.GetConfig("app_name"))
	fmt.Println("OS Name:", helper.GetOsName())
}