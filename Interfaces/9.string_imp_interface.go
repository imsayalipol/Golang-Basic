package main

import "fmt"

type Course interface {
	WebDevelopment()
	EmbeddedDevelopment()
}

func (s student) WebDevelopment() {
	fmt.Println(s, "is a Web Developer")
}

func (s student) EmbeddedDevelopment() {
	fmt.Println(s, "is a Embedded Software Developer")
}

type student string

func main() {

	var c Course

	c = student("John")
	c.WebDevelopment()
	c.EmbeddedDevelopment()

}
