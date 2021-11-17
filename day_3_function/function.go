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

// recursive function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

// defer, panic, recover
func logging() {
	fmt.Println("selesai memanggil logging function")
}

func endApp() {

	fmt.Println("End App")
	message := recover()
	if message != nil {
		fmt.Println("Terjadi ERROR: ", message)
	}

}

func runAppDefer() {
	defer logging()
	// pangil function logging di belakang function!
	//eksekusi meskipun function utama error defer mesti diteteapkan di atas buakn dibawah
	fmt.Println("RUN Application")

	fmt.Println("RUN Application") // tanpa deffer
	logging()
}

func runAppPanic(error bool) {

	defer endApp()
	if error {
		panic("==ERROR==")
	}
	fmt.Println("RUN Application running")
}

func runAppRecover(error bool) {
	defer endApp()
	if error {
		panic("==ERROR==")
	}
	fmt.Println("RUN Application recover")
}

// struct
type Customer struct {
	Name, Address string
	Age           int
	Married       bool
	/*
		struct adalah template data atau template data.

		cara penulisan struct:
		* nama struct
			1 kata = Customer
			2 kata = CustomerData
		* field struct
			kebiasan namanya dengan diawali huruf BESAR
			Nama bukan nama

	*/
}

// struct function/method
// seolah2 struct punya function
func (customer Customer) sayHi(name string) {
	fmt.Println("=== hai! ===", name, "my name is", customer.Name)
}

func (customer Customer) sayHell() {
	fmt.Println("=== hai!  From ===>> ", customer.Name)

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
	fmt.Println("====== recursive ======")
	loop := factorialLoop(5)
	fmt.Println("Loop,", loop)
	fmt.Println("factorial:", 5*4*3*2*1)
	fmt.Println("factorial recursive:", factorialRecursive(5))

	// closure function
	fmt.Println("==== closure ===")
	addr := "ygj"
	counter := 0

	increment := func() {
		// data diluar scope bisa diakses dari dalam scope
		addr = "jkt"
		// addr := "smrg" // tanda deklarasi variabel
		fmt.Println("increment")
		counter++
	}
	// data didalam scope tidak bisa diakses dari luar scope

	increment()
	increment()
	fmt.Println("course: ", counter)
	fmt.Println("addr: ", addr)

	fmt.Println("===== defer, panic, recover ====")
	fmt.Println("=== defer ===")
	runAppDefer()
	fmt.Println("=== panic ===")
	runAppPanic(false)
	fmt.Println("=== recover ===")
	runAppRecover(true)
	fmt.Println("heriipurnama")

	fmt.Println("========= struct =======")
	var heriipurnama Customer
	heriipurnama.Name = "heriipurnama"
	heriipurnama.Address = "Prambanan"
	heriipurnama.Age = 22
	fmt.Println("view all data struct=> ", heriipurnama)
	fmt.Println("nama: ", heriipurnama.Name) // tampil satu2 data
	fmt.Println("Age: ", heriipurnama.Age)   // Age
	// representasi data lebih bagus di struct dari pada map apalagi array!

	// struct Literals
	// mode 1
	joko := Customer{
		Name:    "joko",
		Address: "Ind",
		Age:     23,
	}
	fmt.Println("=====> ", joko)

	// mode 2
	// budi := Customer{"Budi", "jakarta", 25}
	/*
		model diatas rentan Error kalau nama struct ada perubahan
		better mode 1 recomend karena pemagilan struct didefinisikan
		per nilai struct nya.
	*/
	// fmt.Println("=====> ", budi)
	fmt.Println("======== struct method =======")
	heriipurnama.sayHi("Budi")
	heriipurnama.sayHell()
}
