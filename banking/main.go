package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello again"))
		fmt.Fprint(w, "hello world")
	})

	//fmt.Println("server started at localhost:9000")
	http.ListenAndServe("localhost:9000", nil)
}
