package main

import (
	"fmt"
	"sync"
)

type study struct {
	subject    string
	university string
	id         int
}

func main() {

	var wg sync.WaitGroup

	StructCh := make(chan study)
	s := study{
		subject:    "Computer science",
		university: "Boston",
		id:         1122,
	}

	go execute(&wg, StructCh)

	StructCh <- s
}

func execute(wg *sync.WaitGroup, StructCh chan study) {

	details := <-StructCh
	fmt.Println("Details:", details)

	fmt.Println(details.id)
	fmt.Println(details.university)
	fmt.Println(details.subject)
}
