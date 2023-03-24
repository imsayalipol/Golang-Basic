package main

import "fmt"

func main() {
	a := 20
HERE:
	for a < 25 {
		fmt.Println("Hello from Demo func. a =", a)
		fmt.Println("Hello from main func")
		a++
	}
	goto HERE

}
