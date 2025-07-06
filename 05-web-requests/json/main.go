package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("hello")
	person := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println("person data is:", person)

	// convert person into json encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error marshelling", err)
		return
	}
	fmt.Println("person data is : ", string(jsonData))

	// decoding (Unmarshalling)
	var personData Person
	unmarshalErr := json.Unmarshal(jsonData, &personData)
	if unmarshalErr != nil {
		fmt.Println("error unmarshelling", unmarshalErr)
		return
	}
	fmt.Println("person data is:", personData)

}
