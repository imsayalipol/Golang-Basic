package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type employee struct {
	Name string `json:"Name"` //field name here and in json file should be same
	Age  int    `json:"Age"`  // if we write = age int `json:"Age"` it won't return values from json file
}

func main() {

	var emp employee

	file, err := ioutil.ReadFile("employee.json")
	if err != nil {
		log.Fatalf("Some error :", err)
	}

	err = json.Unmarshal([]byte(file), &emp)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("emp Struct: %#v\n", emp)

}
