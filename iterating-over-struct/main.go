package main

import (
	"fmt"
	"reflect"
)

type someStruct struct {
	SomeFeild1 string
	SomeFeild2 interface{}
	SomeFeild3 int
	SomeFeild4 bool
	SomeFeild5 struct {
		hello string
		there interface{}
	}
}

func main() {
	res := someStruct{
		SomeFeild1: "feild1",
		SomeFeild2: 2,
		SomeFeild3: 4,
		SomeFeild4: true,
	}

	s := reflect.ValueOf(&res).Elem()
    fmt.Println(s)
    fmt.Println(s.NumField())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		v := reflect.ValueOf(f.Interface())
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
            fmt.Printf("%v, %v, %v, %v, %v",f,f.Interface(),v,v.Interface(), reflect.Zero(v.Type()).Interface())
		}
	}
}
