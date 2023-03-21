package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println("Hello ,world !")
	var a = 10
	fmt.Println("type: ", reflect.TypeOf(a))
	fmt.Println("size: ", unsafe.Sizeof(a))
}
