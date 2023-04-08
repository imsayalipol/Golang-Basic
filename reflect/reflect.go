package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	Name string
}

func main() {
	// string
	mem := myStruct{"Taylor Swift"}
	t := reflect.TypeOf(mem)
	v := reflect.ValueOf(mem)

	fmt.Println("t: ", t, "v: ", v)

	// another way of checking type:
	var i interface{} = "hello"

	ms, ok := i.(string)
	fmt.Println(ms)
	fmt.Println(ok)
	if ok {
		fmt.Println("i is a string:", mem)
	} else {
		fmt.Println("i is not a string")
	}
	fmt.Println()

	// printing type and value of empty data types
	var a int32
	var f float32
	var b bool
	var arr = [3]int{}
	var sli = []string{}
	var LastName string
	var m = make(map[int]int)
	ptr := new(int)
	type Student struct{}
	var s Student
	type Person interface{}
	var p Person
	var ch chan int
	fmt.Println("a: ", reflect.TypeOf(a), "  Value:", reflect.ValueOf(a))
	fmt.Println("f: ", reflect.TypeOf(f), "  Value:", reflect.ValueOf(f))
	fmt.Println("b: ", reflect.TypeOf(b), "  Value:", reflect.ValueOf(b))
	fmt.Println("arr: ", reflect.TypeOf(arr), "  Value:", reflect.ValueOf(arr))
	fmt.Println("sli: ", reflect.TypeOf(sli), "  Value:", reflect.ValueOf(sli))
	fmt.Println("LastName: ", reflect.TypeOf(LastName), "Value:", reflect.ValueOf(LastName))
	fmt.Println("m: ", reflect.TypeOf(m), "  Value:", reflect.ValueOf(m))
	fmt.Println("ptr: ", reflect.TypeOf(ptr), "  Value:", reflect.ValueOf(ptr))
	fmt.Println("s: ", reflect.TypeOf(s), "  Value:", reflect.ValueOf(s))
	fmt.Println("p: ", reflect.TypeOf(p), "  Value:", reflect.ValueOf(p))
	fmt.Println("ch: ", reflect.TypeOf(ch), "  Value:", reflect.ValueOf(ch))
}
