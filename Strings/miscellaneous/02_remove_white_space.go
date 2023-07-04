package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Sayali Pradeep Satvik"
	family := strings.ReplaceAll(s, " ", "")
	fmt.Println(family)
}
