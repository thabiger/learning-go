package main

import "fmt"

/*
	this time i'm going to use an array instead of a slice here
*/
var input = [8]string{"c", "a", "z", "k", "p", "d", "v", "b"}

func main() {

	fmt.Printf("before:%s, %p\n", input, &input)
	sort(input)
	fmt.Printf("after:%s, %p\n", input, &input)

	/* what has changed that modifications made within the function
	   doesn't reflect on the state of the variables that were passed
	*/

	/*
	   to achieve the same effect as before, function result has to
	   explicitly assigned to the variable
	*/
	input = sort(input)

	fmt.Printf("after assignement:%s, %p\n", input, &input)

}

// just a custom sorting function
func sort(arr [8]string) [8]string {

	/*
		we're operating on the copy of the array within te function.

		changes made within the function would not result
		in changing the passed variable state
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
