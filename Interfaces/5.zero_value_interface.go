package main

import "fmt"

type Person interface {
	info()
}

func main() {
	var p Person
	fmt.Println(p)

}
