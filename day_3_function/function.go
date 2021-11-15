package main

import "fmt"

func sayHelo1() {

	fmt.Println("hello")
}

func sayHeloTo(firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}

// function return value
func getHello(name string) string {
	if name == "" {
		return "hello"
	} else {
		return "Hello" + name
	}

}

// returning multiple values
func getFullName() (string, string) {
	return "herii", "purnama"
}

func main() {
	fmt.Println("========function=======")
	for i := 0; i < 10; i++ {
		sayHelo1()
	}
	sayHeloTo("'heriipurnama'", "'reina'")
	// bisa langsung! //function parameter
	fmt.Println("=======function parameter====")
	firstName := "sofy"
	sayHeloTo(firstName, "purnama")
	sayHeloTo("reina", "sofiana")

	// function return value
	result := getHello("reina")
	fmt.Println("result: ", result)

	// returning multiple values
	fmt.Println("===returning multiple value===")
	firstName, lastName := getFullName()
	fmt.Println("firstName: ", firstName, ",lastName: ", lastName)
	// ignore return values
	fmt.Println("===ignore returning multiple values")
	namaDepan, _ := getFullName()
	fmt.Println("namaDepan: ", namaDepan)
}
