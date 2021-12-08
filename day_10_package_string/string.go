package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("heriipurnama s", "heriipurnama"))
	fmt.Println(strings.Split("heriipurnama s", " "))
	fmt.Println(strings.ToLower("heriipurnama s"))
	fmt.Println(strings.ToUpper("heriipurnama s"))

	fmt.Println(strings.Trim(" heriipurnama ", " "))
	fmt.Println(strings.ReplaceAll("heriipurnama heriipurnama heriipurnama", "heriipurnama", "sofy"))
}
