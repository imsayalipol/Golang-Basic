package main

import "fmt"

// uncomment various ways one at a time and see the output
// way 1
const (
	a = iota
	b
	c
)

//way 2
// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

//way 3
// const (
//     a = iota
//     b
//     c = iota
// )

// way 4
// const (
// 	a = iota	//0
// 	_
// 	b			//2
// 	c			//3
// )

// only in this case counting will restart from 0 for c
// const (
// 	a = iota
// 	b
// )
// const (
// 	c = iota
// )

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
