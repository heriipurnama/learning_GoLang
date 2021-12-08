package database

import "fmt"

var connection string

func init() {
	fmt.Println("function init dipangil!")
	connection = "MYSQL"
}

func GetDatabase() string {
	return connection
}
