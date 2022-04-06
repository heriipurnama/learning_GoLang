/*
   macam-macm go tools:
   command :
       1. go --help
       2. go build
            build package digunakan untuk build package menjadi file executable.
            a. go build
                default nama package
            b. go build -o <nama yang diinginkan>
                go build app
                ** nama hasil build ada di dalam dir package day_19
            c. running file build go
                ** ./<nama file>
                    exp: ./app
        3. go fmt
            untuk format kode golang seperti linternya golang
*/

package main

import "fmt"

func main() {
	fmt.Println("golang go tools !!!")
}
