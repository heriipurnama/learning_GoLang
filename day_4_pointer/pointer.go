package main

import "fmt"

/*
	pointer function di parameter
*/

type Address struct {
	City, Province, Country string
}

type Address12 struct {
	City, Province, Country string
}

func changgeAdressIndonesia(address *Address) { // Dengan Pointer * , untuk menangani banyak value yang ada
	//dalam satu waktu agar tidak copy data dan menghemat memory
	address.Country = "Indonesia"
}

func changgeAdressIndonesia1(address12 Address12) { // Bukan Pointer
	address12.Country = "Indonesia"
}

/*
	pointer function , sebelum function
*/

type Man struct {
	Name string
}

type Female struct {
	Name s9tring
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func (female Female) Married1() {
	female.Name = "Mss." + female.Name
}

func main() {

	/*
	   pas by value
	*/
	// address 1 di copy ke address 2.
	address1 := Address{"subang", "jabar", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println("address1", address1) // address1 tidak berubah
	fmt.Println("address2", address2)

	// pointer
	// pas by reference
	// * tanda pointer
	addr1 := Address{"pekan baru", "aceh", "riau"}
	addr2 := &addr1

	addr2.City = "Jogja"

	fmt.Println("addr1", addr1)
	fmt.Println("addr2", addr2)

	// pointer operator (*)
	add1 := Address{"sleman", "jogja", "Indonesia"}
	add2 := &add1

	/*
	   penulisan ke 2
	*/
	// var address1 Address = Address{"sleman", "jogja", "Indonesia"}
	// var address2 *Address = &address1

	add2.City = "Bandung"

	add2 = &Address{"jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println("addr1=> ", add1)
	fmt.Println("addr2=> ", add2)

	// tanpa operator (*)
	fmt.Println("====== tanpa operator ======")
	adr1 := Address{"sleman", "jogja", "Indonesia"}
	adr2 := &adr1
	adr3 := &adr1

	adr2.City = "Bandung"

	*adr2 = Address{"Kuningan", "DKI Jakarta", "Indonesia"}

	fmt.Println("adr1=> ", adr1)
	fmt.Println("adr2=> ", adr2)
	fmt.Println("adr3=> ", adr3)
	// function new
	fmt.Println("=== function new ====")
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"
	fmt.Println("alamat 1", alamat1)
	fmt.Println("alamat 2", alamat2)
	fmt.Println("=== cara 2 ===")
	var adr4 *Address = new(Address)
	adr4.City = "Aceh"
	fmt.Println("adr4==> ", adr4)

	fmt.Println("=== pointer on function ===")
	address := Address{"Subang", "Jawa Barat", ""}
	address12 := Address12{"Subang", "Jawa Barat", ""}

	changgeAdressIndonesia(&address)
	fmt.Println("===", address, "===")
	// cara kedua
	var addressPointer *Address = &address
	fmt.Println("addressPointer==> ", addressPointer)
	// Bukan pointers
	changgeAdressIndonesia1(address12)
	fmt.Println("=== Not Pointer can Change value -->", address12)
	// Pointer on METHOD
	fmt.Println("==== pointer method ====")
	herii := Man{"Herii."}
	herii.Married()
	fmt.Println("dengan pointer===>> ", herii.Name)

	sofy := Female{"Sofy."}
	sofy.Married1()
	fmt.Println("tdk dengan pointer===>> ", sofy.Name)
}
