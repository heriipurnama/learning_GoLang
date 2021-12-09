package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Contoh struct {
	Name        string
	Description string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}

	}
	return true
}

func main() {
	sample := Sample{"heriipurnama"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0))

	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")

	fmt.Println(required)

	sample.Name = "" // kalau ada isinya maka akan menghasilkan string kosong dan bernilai false!
	fmt.Println(IsValid(sample))
	fmt.Println("------------------")
	contoh := Contoh{"", "b"}
	fmt.Println(IsValid(contoh))
}
