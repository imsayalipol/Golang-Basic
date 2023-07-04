package main

import (
	"fmt"
)

func main() {

	// String Compare
	// comp1 := strings.Compare("a", "A")
	// fmt.Println(comp1)

	// comp2 := strings.Compare("A", "a")
	// fmt.Println(comp2)

	// comp3 := strings.Compare("a", "a")
	// fmt.Println(comp3)

	s := "sayali"
	sli := []rune(s)
	// sli = strings.Split(s, "")
	fmt.Println(sli)

	// var newSli []string
	// for i := len(sli) - 1; i >= 0; i-- {
	// 	newSli = append(newSli, sli[i])
	// }
	// revStr := strings.Join(newSli, "")
	// fmt.Println(revStr)

}
