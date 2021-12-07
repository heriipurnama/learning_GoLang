package helper

import "fmt"

var Application = "learning_GoLang"
var version = 1.0

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("Goodbye", name)
}

/*
   Note: Dalam satu package tidak bisa / tidak boleh, contoh disini package `helper` menggunakan nama fungsi yang sama.
*/
