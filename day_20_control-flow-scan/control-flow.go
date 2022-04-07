package main

import "fmt"

func main() {
	// if <condition> { <consequent> }
	// else

	var y = 16
	if y >= 17 {
		fmt.Printf("Y between 17 and 30\n")
	} else {
		fmt.Printf("Y not between 17 and 30 \n")
	}
	fmt.Println("=================")

	// for <init>; <condition>; <update>;
	// cara 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n", i)
	}
	fmt.Println("=================")

	// cara 2
	a := 0
	for a < 10 {
		fmt.Printf("==> %d \n", a)
		a++
	}
	fmt.Println("=================")

	// switch
	/*
		switch x {
			case 1:
				....
			case 2:
				....
			default:
			    ....
		}
	*/

	var dayOfWeek = 3
	switch dayOfWeek {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	default:
		fmt.Println("Invalid day's")
	}

	fmt.Println("========SWITCH=========")
	switch days := 3; days {
	case 1, 2, 3, 4, 5:
		fmt.Println("weekdays")
	case 6, 7:
		fmt.Println("weekend")
	default:
		fmt.Println("Invalid days")
	}

	// scan
	var x = 21
	switch {
	case x < 18:
		fmt.Println("Kurang dari 18")
	case x >= 18 && x < 25:
		fmt.Println("18 sampai 25")
	case x >= 25 && x < 50:
		fmt.Println("25 sampai 50")
	default:
		fmt.Println("Invalid ")
	}

	// break *hentikkan looping ketika sdh menemukan data yang diinginkan
	for num := 1; num < 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("==> %d \n", num)
			break
		}
	}

	// continue *lanjutkan looping ketika sdh menemukan data yang diinginkan sampai kondisi looping terpenuhi
	for nums := 1; nums < 10; nums++ {
		if nums%2 == 0 {
			continue
		}
		fmt.Printf("%d \n", nums)
	}

	// scan *ambil input dari user dan juga mengabilnpointer sebagai argument
	var appleNum int
	fmt.Printf("Number of Apple? \n")
	fmt.Scan(&appleNum)
	fmt.Print("we have: ", appleNum, "\n")
}
