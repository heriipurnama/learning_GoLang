package main

import "fmt"

func main() {
	// array
	y := [5]int{1, 2, 3, 4}

	for i, y := range y {
		fmt.Println("index : ", i, "value :", y)
	}

	//slice *lebih flexible
	arr := [...]int{1, 2, 3, 4}
	s1 := arr[1:3] // 1 = index, 3 = panjang slice
	s2 := arr[2:4] // 1 = index, 3 = panjang slice

	fmt.Println("S1 :", s1)
	fmt.Println("S2 :", s2)

	// slice 2 argument

	// sli = make([]int, 10) // 2 argumnet

	// sli2 = make([]int, 10, 15) // 3 argumnet

	// buat variabel SLICE dengan make
	s := make([]string, 3)
	fmt.Println("empty: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s) // copy(<tujuan>, <asal>)
	fmt.Println("copy: ", c)
}
