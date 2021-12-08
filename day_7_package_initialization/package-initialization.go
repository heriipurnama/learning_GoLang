package main

import (
	"fmt"
	_ "learning_GoLang/day_7_package_initialization/databases" // Blank Identifier
	database "learning_GoLang/day_7_package_initialization/databases"
)

func main() {
	result := database.GetDatabase()
	fmt.Println("---> ", result)
}
