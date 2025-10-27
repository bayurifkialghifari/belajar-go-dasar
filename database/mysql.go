package database

var database string
var username string
var password string

func init() {
	database = "mysql"
	username = "root"
	password = "password"
}

func GetDatabase() (string, string, string) {
	return database, username, password
}