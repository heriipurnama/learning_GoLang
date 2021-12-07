package main

import (
	"fmt"
	"learning_GoLang/day_6_acces_modifier/helper"
)

func main() {
	// function
	helper.SayHello("heriipurnama")
	// helper.sayGoodBye("heriipurnama") // error tidak bisa diakses

	// variable
	fmt.Println("--> ", helper.Application)
	// fmt.Println("--> ", helper.version) // error tidak bisa diakses
}

/*
   Note :
   Acces Fungsi / variabel diluar package nama, nama Fungsi/Variabel Diawali dengan HURUF BESAR
   jika diawali huruf kecil tidak bisa diakses!
*/
