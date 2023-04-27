package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	fmt.Println("Deleteing elemnts 2,3 from x")
	x = append(x[:1], x[3:]...) // variadic func(unfurling)
	fmt.Println(x)

}
