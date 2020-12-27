package main

import "fmt"

/* in the previous example of working with maps, i use this example:

var person = map[string]string{
	"Name":    "John",
	"Surname": "Doe",
	"Age":     "29",
}

as a map can only contain values of one type, it requires to accept
the inconvenience of storing age as a string imstead of integer. this
limitation can be easily overcome with structs

*/

type person struct {
	Name    string
	Surname string
	Age     int
}

func main() {

	john := person{
		Name:    "John",
		Surname: "Doe",
		Age:     29,
	}

	fmt.Printf("before: %v %v, Age: %v, %p\n\n", john.Name, john.Surname, john.Age, &john)

	modify_age(john, 64)

	fmt.Printf("after: %v %v, Age: %v, %p\n\n", john.Name, john.Surname, john.Age, &john)

	/*
		structs arguments are passed by values - changes made within the function
		are not visible to the caller
	*/

	john = modify_age(john, 64)

	fmt.Printf("after assignement: %v %v, Age: %v, %p\n\n", john.Name, john.Surname, john.Age, &john)

	/*
		we have to assign the result of the function to the variable to make the change effective
	*/

}

func modify_age(p person, age int) person {
	fmt.Printf("within the function:%s, %p\n\n", p, &p)

	p.Age = age

	return p
}
