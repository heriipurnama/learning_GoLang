package main

import "fmt"

func sayHelo1() {

	fmt.Println("hello")
}

func main() {
	fmt.Println("========function=======")
	for i := 0; i < 10; i++ {
		sayHelo1()
	}

}
