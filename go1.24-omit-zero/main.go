package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Pet  Pet `json:",omitempty,omitzero"`
}

type Pet struct {
	Name string
	Age  int
}

func displayJsonFormat(person *Person) {
	b, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(b))
}

func main() {
	person := &Person{
		Name: "John",
		Age:  30,
		Pet:  Pet{Name: "Max", Age: 5},
	}
	displayJsonFormat(person)

	person = &Person{
		Name: "John",
		Age:  30,
		Pet:  Pet{},
	}
	displayJsonFormat(person)
}
