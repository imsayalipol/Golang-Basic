package main

import "fmt"

type maths struct{}

func (m *maths) add(numbers ...int) int {

	add := 0

	for _, n := range numbers {
		add += n
	}

	return add
}

func main() {

	m := &maths{}

	fmt.Println(m.add(1, 2, 3))
	fmt.Println(m.add(100, 200, 400, 900, 500))
}
