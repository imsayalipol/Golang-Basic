package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Panic Code")

	panic("problem")
	fmt.Println("We on 12")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	fmt.Println("We afe checking here")
}
