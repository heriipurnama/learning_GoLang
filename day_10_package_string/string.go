package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("heriipurnama s", "heriipurnama"))
	fmt.Println(strings.Split("heriipurnama s", " "))
	fmt.Println(strings.ToLower("heriipurnama s"))
	fmt.Println(strings.ToUpper("heriipurnama s"))

	fmt.Println(strings.Trim(" heriipurnama ", " "))
	fmt.Println(strings.ReplaceAll("heriipurnama heriipurnama heriipurnama", "heriipurnama", "retno"))

	//compare { return decimal }
	//compare(a,b)
	//0 if a==b, -1 if a < b, +1 if a>b

	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println("-----------------------------------")

	//contains { return boolean }
	//contains(s, substr)
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println("-----------------------------------")

	//HasPrefix
	//HasPrefix(s, prefix)
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", "G"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println("-----------------------------------")

	// index
	// index(s, substr)
	fmt.Println("Index")
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	fmt.Println("-----------------------------------")

	// to Upper
	fmt.Println("To Upper")
	fmt.Println(strings.ToUpper("kode"))
	fmt.Println(strings.ToLower("KODE"))
	fmt.Println("-----------------------------------")

	//TrimSpace
	// \t : space character
	// \n : new line
	fmt.Println("trimSpace")
	fmt.Println(strings.TrimSpace("\t\n Kode Trim \n"))
	fmt.Println("-----------------------------------")

	// Atoi = conversi dr string s kedalam integer.
	v := "10"
	vi := 10
	fmt.Println(strconv.Atoi(v))
	fmt.Println("-----------------------------------")

	//ItoA
	//formatFloat
	//ParseFloat
	fmt.Println(strconv.Itoa(vi))

}