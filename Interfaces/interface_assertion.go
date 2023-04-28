package main

import "fmt"

type human interface {
	speak()
}

type man struct {
	name string
}

type woman struct {
	fname string
}

func (m man) speak() {
	fmt.Println("Im ", m.name)
}

func (w woman) speak() {
	fmt.Println("Im ", w.fname)
}

// type assertion
func bar(h human) {
	fmt.Println("I was passed to bar", h)
	switch h.(type) {
	case man:
		fmt.Println("Hi from bar, I am ", h.(man).name)

	case woman:
		fmt.Println("Hi from bar, I am ", h.(woman).fname)
	}
}

func main() {
	man := man{"John"}
	man.speak()

	wom := woman{"Sofia"}
	wom.speak()

	bar(man)
	bar(wom)
}
