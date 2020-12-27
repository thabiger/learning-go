package main

import "fmt"

var person = map[string]string{
	"Name":    "John",
	"Surname": "Doe",
	"Age":     "29",
}

func main() {

	fmt.Printf("before:%s, %p\n\n", person, &person)

	for k, v := range person {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	modify(person, "Name", "Jane")
	fmt.Printf("before:%s, %p\n\n", person, &person)
}

func modify(person map[string]string, k string, v string) map[string]string {
	fmt.Printf("within the function:%s, %p\n\n", person, &person)

	person[k] = v

	return person
}
