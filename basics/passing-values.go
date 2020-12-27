package main

import (
	"fmt"
	"strings"
)

var (
	s = "default value"
)

func main() {

	fmt.Println("-- Passing by value example --")
	fmt.Println("value outside the function, before invocation: ", s, "addr: ", &s)

	// argument passed by the value
	uppercaseByValue(s)

	// original value remains untouched
	fmt.Println("value outside the function, after invocation: ", s, "addr: ", &s)

	fmt.Println("\n\n-- Passing by pointer example --")
	uppercaseByPtr(&s)

	// original value remains untouched
	fmt.Println("value outside the function, after invocation: ", s, "addr: ", &s)

}

func uppercaseByValue(s string) {
	fmt.Println("value within the function, before change: ", s, "addr: ", &s)
	/*
		although we're changing the value, we're operating on the copy
		so it should remain untouched outside of the function
	*/
	s = strings.ToUpper(s)

	fmt.Println("value within the function, after change: ", s, "addr: ", &s)
}

func uppercaseByPtr(s *string) {
	fmt.Println("value within the function, before change: ", *s, "addr: ", &s)
	/*
		this time we're operating directly on the passed variable, so
		changes made here should be instantly visible outside the function
	*/
	*s = strings.ToUpper(*s)

	fmt.Println("value within the function, after change: ", *s, "addr: ", &s)
}
