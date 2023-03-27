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
