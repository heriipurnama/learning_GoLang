package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("-- Argument : ")
	fmt.Println(args)

	fmt.Println("get argument-> ", args[0]) // data ke [0] lokasi file Binary
	fmt.Println("get argument-> ", args[1]) // data Argument

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname-> ", name)
	} else {
		fmt.Println("Error-> ", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
	/*
	   Note : Jalankan Perintah
	   export saat di terminal sebelum eksekusi program
	   os.Getenv,
	   os.Getenv digunakan untuk memasukkan data .env kedalam program, atau
	   setting .env pada program
	*/
}
