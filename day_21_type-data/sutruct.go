package main

import "fmt"

// struct digunakan untuk pengelompokkan data

type person struct {
	name string
	age  int64
}

func main() {
	fmt.Println("person", person{"raka", 30})
	fmt.Println("person", person{name: "rakaK", age: 30})
	fmt.Println("person", person{name: "dean"})
	fmt.Println("person", &person{name: "andi", age: 30})

	s := person{name: "Putra", age: 23}

	fmt.Println("name: ", s.name)

	sp := &s
	fmt.Println("age :", sp.age)

}
