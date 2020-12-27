package main

import (
	"fmt"
	"strings"
)

func main() {

	printTree()

	fmt.Print("Your gifts: ")
	/*
	  the first argument returned by range executed on a list
	  is an index, that we're ignoring here
	*/
	for _, gift := range gifts(4) {
		fmt.Print(gift, " ")
	}
}

func gifts(value int) []string {
	// maps stores gifts and their coresponding sizes
	giftsList := map[string]int{"dollhouse": 4, "car": 1, "doll": 1, "book": 1, "bike": 3, "scooter": 3}

	give := []string{}

	// it's not a fair algorithm, but show how 'for loops' work :)
	for k, v := range giftsList {
		value -= v
		if value < 0 {
			break
		}
		give = append(give, k)
	}

	return give
}

func printTree() {
	// multiple assigne statement allows to initialize more then one variables
	for size, line := 20, 0; line <= size; line++ {
		if line%2 == 0 {
			continue
		}
		fmt.Println(strings.Repeat(" ", size-line), strings.Repeat("*", line))
	}
}
