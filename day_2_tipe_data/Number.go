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
	// single variable

	const tool string = "hammers"
	const nature = "earth"

	// err error
	fmt.Println("tool : ", tool)
	fmt.Println("nature : ", nature)

	// multiple variable
	const (
		price int32 = 10
		sat         = "pcs"
	)

	fmt.Println("price :", price)
	fmt.Println("sat : ", sat)

	/**
	 * Konversi Tipe Data
	 */

	//3278
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)

	var nilai8 int8 = int8(nilai32)

	fmt.Println("nilai32", nilai32)
	fmt.Println("nilai64", nilai64)
	fmt.Println("nilai8", nilai8)

	/**
	 * Konversi kode program tipe data(2)
	 */

	var names = "heriipurnama"
	var e = name[0]         // dipersingkat penulisan aslinya e byte = name[1]
	var eString = string(e) // konversi dari byte/uint8 ke string

	/*
				note :
		        byte alias uint8
	*/

	fmt.Println("name : ", names)
	fmt.Println("eString : ", eString)

	/*
	   type Declarations / tipe data alias / membuat tipe data baru dari tipe data yg sudah ada.
	*/

	type NoKTP string
	type married bool

	var ktpHerii NoKTP = "99999999"
	var marriedStatus married = true

	fmt.Println("NoKTP : ", ktpHerii)
	fmt.Println("new Number : ", NoKTP("888888888"))
	fmt.Println("status : ", marriedStatus)

	/*
	   operasi aritmatika
	*/

	var sum = 10 + 10
	var a = 10
	var b = 10
	var c = 10

	var rest = a + b
	var result = a * c

	fmt.Println("sum", sum)
	fmt.Println("rest", rest)
	fmt.Println("result", result)

	// augmented assigment

	var i = 10
	i += 10

	fmt.Println("hasiil : ", i)

	// operasi perbandingan
	var name1 = "heri"
	var name2 = "heri"

	var resultName bool = name1 == name2

	fmt.Println("resultName : ", resultName)

	var val1 = 100
	var val2 = 200

	fmt.Println("val : ", val1 > val2)
	fmt.Println("val : ", val1 < val2)
	fmt.Println("val : ", val1 == val2)
	fmt.Println("val : ", val1 != val2)

	// operasi boolean

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println("============================")
	fmt.Println("lulus : ", lulus)

	// opretator

	var ujian = 80
	var absensi1 = 75

	// cara 1
	var lulusUjian = ujian >= 80
	var lulusAbsensi1 = absensi1 >= 80

	var lulus1 = lulusUjian && lulusAbsensi

	fmt.Println("============================")
	fmt.Println("lulus : ", lulusUjian)
	fmt.Println("absensi : ", lulusAbsensi1)
	fmt.Println("lulus : ", lulus1)

	// cara 2 singkat

	fmt.Println("cara 2 : ", ujian >= 80 && absensi1 >= 80)

	// Array
	var namesArray [3]string
	namesArray[0] = "herii"
	namesArray[1] = "purnama"
	namesArray[2] = "ada lagi"

	fmt.Println("names : ", namesArray[0])
	fmt.Println("names : ", namesArray[1])
	fmt.Println("names : ", namesArray[2])

	var vals1 = [3]int{
		1,
		2,
		3,
	}

	// mengubah isi data array
	vals1[2] = 10
	fmt.Println("vals1 : ", vals1)
	fmt.Println("vals1-0 : ", vals1[0])
	fmt.Println("vals1-2 : ", vals1[2])

	// cek panjang array
	var lagi [5]string

	fmt.Println("length : ", len(vals1))
	fmt.Println("length_lagi : ", len(lagi))

	// Slice
	namess := [...]string{"herii", "purnama", "budi", "doni", "dedi"}
	slices := namess[0:3] // array[low:high]
	slices2 := namess[3:]

	fmt.Println("datas array", slices)
	fmt.Println("slices0 : ", slices[0])
	fmt.Println("slices1 : ", slices[2])
	fmt.Println("length array : ", len(slices))

	fmt.Println("cap array : ", cap(slices))
	fmt.Println("slice2 : ", slices2)

	// function make slices

	newSlice := make([]string, 2, 5)
	newSlice[0] = "herii"
	newSlice[1] = "purnama"

	fmt.Println("newSlice: ", newSlice)
	fmt.Println("lengthSlice: ", len(newSlice))
	fmt.Println("capSlice: ", cap(newSlice))

	// copy slices newSlice ke copySlice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println("copySlice: ", copySlice)

	// perbedaan array dan slice
	iniArray := [...]int{0, 1, 2, 3, 4}
	iniArrays := [5]int{0, 1, 2, 3, 4}

	iniSlice := []int{0, 1, 2, 3, 4}

	fmt.Println("iniArray: ", iniArray)
	fmt.Println("iniArrays: ", iniArrays)
	fmt.Println("iniSlice: ", iniSlice)

	// slice append
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println("days:", days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println("daySlice:", daySlice2)
	fmt.Println("days:", days)

	// function make

	newSlice1 := make([]string, 2, 5) // slice sdh didefinisikan
	newSlice1[0] = "herii"
	newSlice1[1] = "purnama"

	fmt.Println("newSlice", newSlice1)
	fmt.Println("newSlice len", len(newSlice1))
	fmt.Println("newSlice cap", cap(newSlice1))

	// copy slice
	fmt.Println("-copy slice-")
	fromSlice := days[2:4]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("fromSlice len:", len(fromSlice))

	copySlice1 := make([]string, len(newSlice1), cap(newSlice1))

	copy(copySlice1, newSlice1)
	fmt.Println("copySlice :", copySlice1)

	fmt.Println("=========TIPE DATA slice==========")
	iniArray1 := [4]int{1, 2, 3, 4}
	iniSlice1 := []int{1, 2, 3, 4}

	fmt.Println("ini aray", iniArray1)
	fmt.Println("ini slice", iniSlice1)

	fmt.Println("=== tipe data map")
	person := map[string]string{
		"name":    "herii",
		"address": "jgj",
	}
	// nambah data
	person["title"] = "entrepreneur"

	fmt.Println("===")
	fmt.Println("person", person)
	fmt.Println("person name: ", person["name"])
	fmt.Println("person addres: ", person["address"])
	fmt.Println("person title: ", person["title"])

	// make
	book := make(map[string]string)

	book["title"] = "Golang"
	book["author"] = "heriipurnama"
	book["page"] = "126"
	fmt.Println("sebelum dihapus==> ", book)
	delete(book, "page")
	fmt.Println("sesudah Book : ", book)

	// if statemen
	nameEmployee := "herii"

	if nameEmployee == "herii" {
		fmt.Println("hello herii")
	} else if nameEmployee == "purnama" {
		fmt.Println("hai purnama")
	} else {
		fmt.Println("HI, boleh kenalan")
	}

	// if short statement
	// if standart

	/*
		var length = len(name)
		if length > 5 {
			fmt.Println("nama terlau panjang")
		} else {
			fmt.Println("nama sudh benar")
		}
	*/

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlau panjang")
	} else {
		fmt.Println("nama sudh benar")
	}
	// switch statement

	alamat := "Prambanan"

	switch alamat {
	case "Prambanan":
		fmt.Println(" case oke")
	case "Klaten":
		fmt.Println("salah")
	default:
		fmt.Println("salah")
	}

	// short statement switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalau panjang")
	case false:
		fmt.Println("nama sudah benar ")
	}
	// switch tanda kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
	// for loop on
	fmt.Println("==========FOR=========")
	counter := 1
	for counter <= 10 {
		fmt.Println("perulangan ke-", counter)
		counter++
	}
	// for statement
	fmt.Println("==========FOR STATEMENT=========")
	// init statement ,					// post statement
	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("perulangan ke-", counter1)
	}
	// for range statement
	// acces data secara manual
	// string iterasi
	fmt.Println("========For Iterasi========")
	sliceName := []string{"herii", "purnama", "reina"}
	for i := 0; i < len(sliceName); i++ {
		fmt.Println("perulangan", sliceName[i])
	}
	// for range range statement
	fmt.Println("========For range=======")
	for index, sliceName := range sliceName {
		fmt.Println("index", index, "=", sliceName)
		fmt.Println("value", sliceName)
	}

	// map
	person1 := make(map[string]string)
	person1["name"] = "heriipurnama"
	person1["tittle"] = "dev."

	for key, value := range person1 {
		fmt.Println(key, "=", value)
	}

	// break into
	fmt.Println("========break=======")
	for i = 0; i < 10; i++ {

		fmt.Println("perulangan", i)

		if i == 5 {
			break
		}

	}
	fmt.Println("========continue=======")
	for i = 0; i < 10; i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan", i)

	}
	fmt.Println("========function=======")
    
}
