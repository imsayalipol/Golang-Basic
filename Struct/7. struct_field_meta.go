package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Education string `json:"edu"`
	Employed  bool   `json:"emp"`
}

func main() {

	p := Person{
		Name:      "John",
		Age:       28,
		Education: "BE",
		Employed:  true,
	}

	marshalPerson, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(marshalPerson))
}
