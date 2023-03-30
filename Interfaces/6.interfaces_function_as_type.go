package main

import "fmt"

type square struct {
	width int
}
type rectangle struct {
	width  int
	length int
}

// circle implementing the interface
func (s square) Area() int {
	return s.width * s.width
}
func (s square) Perimeter() int {
	return 2 * s.width
}

// rectangle implementing the interface
func (r rectangle) Area() int {
	return 2*r.length + 2*r.width
}
func (r rectangle) Perimeter() int {
	return 2 * r.width * r.length
}

type shapes interface {
	Area() int
	Perimeter() int
}

func main() {
	var s []shapes //interface type slice

	squ := &square{width: 10}
	s = append(s, squ)
	rec := rectangle{width: 5, length: 10}
	s = append(s, rec)

	for _, shape := range s {
		fmt.Println("area: ", shape.Area(), "perimeter:", shape.Perimeter())
	}
}
