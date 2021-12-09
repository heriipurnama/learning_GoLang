package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()
	fmt.Println("now: ", now)
	fmt.Println("now: ", now.Year())
	fmt.Println("now: ", now.Month())
	fmt.Println("now: ", now.Day())

	fmt.Println("now: ", now.Hour())
	fmt.Println("-----------------------------------")
	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println("utc: ", utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println("parse: ", parse)
}
