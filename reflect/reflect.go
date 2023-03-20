package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	Name string
}

func main() {
	s := myStruct{"Taylor Swift"}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	fmt.Println("t: ", t, "v: ", v)

	// another way of checking type:
	var i interface{} = "hello"

	p, ok := i.(string)
	fmt.Println(p)
	fmt.Println(ok)
	if ok {
		fmt.Println("i is a string:", s)
	} else {
		fmt.Println("i is not a string")
	}

}
