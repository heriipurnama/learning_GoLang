package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("herii")
	data.PushBack("purnama")

	data.PushBack("heriipurnama") // menambahkan ke plg belakang!
	data.PushBack("retno")
	data.PushBack("sari")

	data.PushFront("depan-> sari") // menambahkan ke plg kedepan

	fmt.Println("Front -> ", data.Front().Value) // ambil dari depan
	fmt.Println("Back -> ", data.Back().Value)   // ambil dari belakang

	fmt.Println("next -> ", data.Back().Next()) // ambil dari data setelahnya
	// cth. Next().Next() atau Prev().Prev()
	fmt.Println("====*ITERASI*====")
	fmt.Println("============================")
	// iterasi untuk mengambil semua data diLIST
	// dari depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println("e.value-> ", e.Value)
	}
	fmt.Println("============================")
	// dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println("e.value-> ", e.Value)
	}

}
