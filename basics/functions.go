package main

import "fmt"

var input = []string{"c", "a", "z", "k", "p", "d", "v", "b"}

func main() {

	fmt.Printf("before:%s, %p\n", input, &input)
	sort(input)
	fmt.Printf("after:%s, %p\n", input, &input)

	/*
		what's interesting here that the slice is not passed by the value,
		as primitives was. the reason for that is that a slice is actually
		a reference to the array that it is builded on.

		the effect is that changes made within the function reflects on
		the state of the passed variable.
	*/
}

// just a custom sorting function
func sort(arr []string) []string {

	/*
		although the slice points to the different address in memory,
		it has the same underlying array

		therefore the changes made to the array by the slice
		will result also in changes seen thorugh the 'input' variable
		that was passed to the function
	*/
	fmt.Printf("argument addr within the function: %p\n", &arr)

	for o := len(arr) - 1; o > 0; o-- {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				t := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = t
			}
		}
	}

	return arr
}
