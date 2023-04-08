package main

import (
	"fmt"
	"reflect"
)

type Person interface{}

func main() {
	var p Person
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p)

}
