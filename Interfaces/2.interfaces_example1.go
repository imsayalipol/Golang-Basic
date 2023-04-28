package main

import "fmt"

type Shape interface {
	Area() int
	// Perimeter() int
}

type Rectangle struct {
	width  int
	height int
}

type Square struct {
	length int
}

func (r *Rectangle) Area() int {
	return 2 * r.width * r.height
}

func (s Square) Area() int {
	return s.length * s.length
}

func geometry(s Shape) {
	fmt.Println("s : ", s)
	fmt.Println(s.Area())
	// fmt.Println(s.Perimeter())
}

func main() {
	r := Rectangle{width: 10, height: 4}
	geometry(&r) // when receiver is receiving a pointer value we pass it as &r

	s := Square{length: 10}
	geometry(s) // when receiver is receiving as value we pass it as s

}
