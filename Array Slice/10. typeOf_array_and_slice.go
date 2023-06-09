package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	// Use reflect.TypeOf to get the type of a variable.
	arrType := reflect.TypeOf(arr)
	sliceType := reflect.TypeOf(slice)

	fmt.Println("Reflect TypeOf ===", arrType)
	fmt.Println("Reflect TypeOf sliceType ===", sliceType)

	fmt.Println("Kind arr ==>", arrType.Kind())
	fmt.Println("Kind slice==>", sliceType.Kind())
}
