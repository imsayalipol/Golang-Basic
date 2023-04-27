package main

import "fmt"

func main() {

	//arrays
	a := [5]int{1, 2, 3, 4, 5} // create
	fmt.Println(a)
	fmt.Printf("%p\n", &a) // address of a
	a[3] = 10              // update
	fmt.Println(a)
	b := append(a[:1], a[3:]...) // delete, array can't be resized and elements can't be deleted
	fmt.Println("new array b:", b)
	// fmt.Println(a[10]) // out of bount error

	fmt.Println()

	// slice
	xi := make([]int, 7) // create
	fmt.Println(xi)
	xi[0] = 10
	xi[1] = 11
	xi[2] = 12
	xi[3] = 13
	xi[4] = 14
	xi[5] = 15

	fmt.Println(xi) // adding elements
	xi[4] = 40      // updating
	fmt.Println(xi)
	xi = append(xi[:4], xi[5:]...) // deleteing 4th element
	fmt.Println(xi)

	// fmt.Println(xi[10]) // reading non existing index will through error

	// map
	m := make(map[int]string) // create
	fmt.Println(m)

	m[1] = "Taylor"
	m[2] = "Dua"
	m[3] = "Ariana"

	fmt.Println(m)

	m[2] = "Selena" // update
	fmt.Println(m)

	delete(m, 2) // delete an element
	v, ok := m[2]
	fmt.Println(m, v, ok) // check deleted value
}
