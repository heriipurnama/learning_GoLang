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

// function value
func getGoodBye(name string) string {
	return "Good Bye" + name
}

//function as parameter
//type declalation
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	resultFilter := filter(name)
	fmt.Println("hello ", resultFilter)
}

// tanpa type declalation

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	resultFilter := filter(name)
// 	fmt.Println("hello ", resultFilter)
// }

func spamFilter(name string) string {
	if name == "jangkrik" {
		return "xxx"
	} else {
		return name
	}
}

// anonymous function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked!", name)
	} else {
		fmt.Println("welcome", name)
	}
}

// func blacklist(name string) bool {
// 	return name == "admin"
// }

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

	fmt.Println("========= function values =======")
	goodbye := getGoodBye
	fmt.Println("goodbye: ", goodbye(" love"))
	//function as parameter
	sayHelloWithFilter("heri", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("jangkrik", filter)
	// anonymous function
	fmt.Println("====Anonymous function===")
	// function disimpan di variabel
	blacklist := func(name string) bool {
		return name == "anjing!"
	}
	registerUser("heri", blacklist)
	registerUser("anjing!", blacklist)
	// function langsung ditulis di parameter
	registerUser("heriipurnama", func(name string) bool {
		return name == "root"
	})
}
