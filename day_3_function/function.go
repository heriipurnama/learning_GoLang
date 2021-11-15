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

// named return values
func getComplete() (firstName, middleName, lastName string) {
	firstName = "herii"
	middleName = "reina"
	lastName = "purnama"

	return
	// return firstName, middleName, lastName // type explicit
}

// variadic function
func sumAll(number ...int) int { // variadic function
	total := 0
	for _, number := range number {
		total += number
	}
	return total
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
	// named return values
	fmt.Println("===named return values====")
	firstName, middleName, lastName := getComplete()
	namaDepan, namaTengah, namaBelakang := getComplete()
	fmt.Println("firstName: ", firstName, ",middleName: ", middleName, ",lastName: ", lastName)
	fmt.Println("namaDepan: ", namaDepan, ",namaTengah: ", namaTengah, ",namaBelakang: ", namaBelakang)

	fmt.Println("===varidic function =======")
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println("total: ", total)
	// slice slice
	numbers := []int{10, 10, 10, 10, 10, 5}
	total = sumAll(numbers...)
	fmt.Println("total: ", total)

}
