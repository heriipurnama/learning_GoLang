package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")
	/*
			   NOTE :
			   "e([a-z])o" --> depan nya hufur `e` belakangnya huruf `o`
		       dan ditengah huruf kecil
	*/

	fmt.Println("--> ", regex.MatchString(("eko")))
	fmt.Println("--> ", regex.MatchString(("eto")))
	fmt.Println("--> ", regex.MatchString(("eDo")))

	search := regex.FindAllString("eko eka edo eto eta eki", 1)
	search2 := regex.FindAllString("eko eka edo eto eta eki", -1)
	fmt.Println("search--> ", search)   //  cari di parameter, parameter = `1`
	fmt.Println("search2--> ", search2) //  jk parameter = `-1`, akan dicocokan dg semua regex yang ada di `regex` dg inputanya
}
