// Serialization is the process of converting an object into a stream
// of bytes or a string representation that can be transmitted or stored.
// In this case, the json.Marshal function serializes the BostonUniversity struct
// into a JSON string representation that can be transmitted over a network or stored in a file.
// This is a common technique used in distributed systems where applications need to communicate
// with each other over a network.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type University struct {
	Name         string
	StudentCount int
	Courses      []string
}

func main() {

	var BostonUniversity = University{
		Name:         "Boston",
		StudentCount: 500,
		Courses:      []string{"AI", "ML", "Data Science", "Embedded Systems"},
	}

	// Marshalling
	MarshalData, err := json.Marshal(BostonUniversity)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Marshaled struct is : %s\n", MarshalData)

	MarshalIndentData, err := json.MarshalIndent(BostonUniversity, "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Marshal Indent is : %s", MarshalIndentData)

	var newUni University

	// Unmarshaling
	err = json.Unmarshal(MarshalData, &newUni)
	fmt.Println("newUni :", newUni)
}
