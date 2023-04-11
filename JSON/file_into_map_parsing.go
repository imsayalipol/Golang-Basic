package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	a := make(map[int]string)

	file, err := ioutil.ReadFile("temp.json")
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	json.Unmarshal([]byte(file), &a)

	fmt.Println(a)
}
