package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true") // nilai hanya bisa true/false selain itu Error.

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64) // hanya bisa dimasukkan number jika dimasukkan String akan Error.

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("error", err.Error())
	}
}
