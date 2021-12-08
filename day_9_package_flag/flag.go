package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "127.0.0.1", "Put your HOST")
	var user *string = flag.String("user", "root", "Put your USER")
	var password *string = flag.String("password", "123!", "Put your PASSWORD")
	var age *int = flag.Int("age", 18, "Put your AGE!")

	flag.Parse() // untuk pemaggilan data yang CUSTOM, supdaya data yang dimasukkan DINAMIS bukan DEFAULT lagi!

	fmt.Println("host", host) // ketika eksekusi POINTER tanpa tanda pointer *
	fmt.Println("user", user)
	fmt.Println("password", password)
	fmt.Println("age", age)
	fmt.Println("----------------")
	fmt.Println("host", *host) // ketika eksekusi POINTER dengan tanda pointer *
	fmt.Println("user", *user)
	fmt.Println("age", *age)
	// eksekusi
	// command default : go run flag.go
	// command custom : go run flag.go -host=localhost -user=root -password=123 -age=20
	// command custom Error jika salah memasukkan data yang tidak sama dengan type data : go run flag.go -host=localhost -user=root -password=123 -age=herii
	/* RESULT :
	    -age int
	      Put your AGE! (default 18)
	-host string
	      Put your HOST (default "127.0.0.1")
	-password string
	      Put your PASSWORD (default "123!")
	-user string
	      Put you
	*/
}
