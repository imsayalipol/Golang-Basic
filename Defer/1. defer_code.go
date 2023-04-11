package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	err := WriteSomeText("SOme Text")

	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Text added to file successfully")
}

func WriteSomeText(text string) error {
	file, err := os.Open("temp.txt")

	if err != nil {
		return err
	}

	defer file.Close()

	n, err := file.WriteString("SOme tExt addeD")
	if err != nil {
		fmt.Println("Error occured while writing into file\n")
		return err
	}

	fmt.Printf("No of bytes written are :%d", n)
	file.Close()
	return nil
}
