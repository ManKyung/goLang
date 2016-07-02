// example.go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"정만경"`
	age  int    `tag1:"나이" tag2:"29"`
}

func main() {
	p := Person{}
	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	var a *int = new(int)
	*a = 1

	fmt.Println(*a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(*a))
	fmt.Println(reflect.ValueOf(a).Elem().Int())
}
