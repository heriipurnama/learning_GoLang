package main

import "fmt"

func main() {
	/* *
	 * Tipe Data NUmber
	 */
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)

	/* *
	 * Tipe Data Boolean
	 */
	fmt.Println("Benar", true)
	fmt.Println("Salah", false)

	/* *
	 * Tipe Data string
	 */
	fmt.Println("Benar")
	fmt.Println("salah")

	fmt.Println("Count string word ", len("Benar"))
	fmt.Println("Benar"[0]) // get value on index 0 on string Benar, view byte on index Benar not B.

	/* *
	 *  Variable
	 */
	var name string
	name = "heriipurnama"
	fmt.Println(name)

	name = "herpur"
	fmt.Println(name)

	var firstName = "Angga" // jika harcode value, maka tidak perlu menyebutkan tipe datanya.
	fmt.Println(firstName)

	var age = 28
	fmt.Println(age)

	lastName := "--Purnama--" // tidak dengan kata kunci var. tidak vajib tapi hanya untuk deklari awal
	fmt.Println(lastName)

	// multiple variable
	var (
		address = 90
		jk      = "Man"
	)

	fmt.Println("age : ", address)
	fmt.Println("jk : ", jk)

	/**
	 * constant variabel
	 */
	fmt.Println("========CONST===========")
	const tool string = "hammers"
	const nature = "earth"

	// err error
	fmt.Println("tool : ", tool)
	fmt.Println("nature : ", nature)
}
