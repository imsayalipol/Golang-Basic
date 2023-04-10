package main

import "fmt"

type Books struct {
	Name  string
	Pages int
	Price float32
}

func test(b *Books) {
	b.Pages = 555
}

func main() {

	book := Books{
		Name:  "The Go Programming Language",
		Pages: 300,
		Price: 500.55,
	}

	// way 1
	bptr := &book

	fmt.Println("bptr is &bptr:", bptr)
	fmt.Println("bptr is *bptr:", *bptr)

	// accessing struct fields with the pointer

	fmt.Println("Name of var book :", (*bptr).Name)
	fmt.Println("Price of var book :", bptr.Price)

	//way 2  this points to empty struct
	sptr := new(Books)
	fmt.Println("*sptr :", *sptr)
	fmt.Printf("Another way of printing : %+v\n", *sptr) // prints keys:values

	var book2 = Books{
		Name:  "C programming",
		Pages: 300,
		Price: 440,
	}
	test(&book2)
	fmt.Println("new changes in book2:", book2)
}
