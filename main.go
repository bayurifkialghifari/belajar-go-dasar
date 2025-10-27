package main

import (
	"errors"
	"fmt"
	"golang-package/database"
	"golang-package/helper"
	_ "golang-package/unused"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(helper.GetConfig("app_name"))
	fmt.Println("OS Name:", helper.GetOsName())

	// Get database configuration from mysql package
	db, user, pass := database.GetDatabase()
	fmt.Println("Database:", db)
	fmt.Println("Username:", user)
	fmt.Println("Password:", pass)

	// Example usage of Bagi function
	result, err := Bagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func Bagi(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return a / b, nil
}