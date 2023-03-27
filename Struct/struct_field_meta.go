package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name      string `json:"n"`
	Age       int    `json:"a"`
	Education string `json:"ed"`
	Employed  bool   `json:"em"`
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
