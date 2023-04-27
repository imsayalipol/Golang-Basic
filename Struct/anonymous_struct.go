package main

import "fmt"

func main() {
	// struct doesn't have a name
	m := struct {
		name string
		age  int
	}{
		"James Bond",
		32,
	}
	fmt.Println(m)
}
