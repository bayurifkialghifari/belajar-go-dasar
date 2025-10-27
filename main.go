package main

import (
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
}