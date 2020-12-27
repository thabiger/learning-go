package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("-- First attempt with faulty arg --")
	test("")
	fmt.Println("-- Second one with a valid one --")
	test("foo bar")

}

func test(arg string) {
	_, err := example(arg)

	if err != nil { //values of returned errors should be checked
		fmt.Println("An error occured", err)
	} else {
		fmt.Println("All good!")
	}
}

func example(arg string) (result string, err error) {

	if arg == "" {
		err = errors.New("Input cannot be empty!")
	}

	return arg, err
}
