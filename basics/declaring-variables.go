package main //executable (not a library)

import (
	"fmt"
	"reflect"
)

// simple variables initialization
var i int

// simple variables initialization and declaration
var j = 0

/*
  - multiple variables can be initialized in one block
  - package-level variables are available to all functions within the package -
    they are global in scope
*/
var (
	/*
		package-level non-declaration initializing statement
		will result in numeric values initialized with zeroes
		and strings as empty
	*/
	num int64
	str string
	flt float32

	/*
		it is also posible to infer types - go will guess
		the type from the declared value
	*/
	infNum = 1
	infStr = "str"
	infFlt = 1.1

	/*
		variable of one particular type can be converted
		by casting to a variable of antoher type

		in this example float is converted to type int
	*/
	castedNum = int(infFlt)
)

const (
	constStr = "Constant string"
	constNum = 6
	constFlt = 1.1
)

func main() {

	/*
		- this variable has a local scope to the main function
		- ":=" declaration works only within the functions body
		- variables declared locally must be used, if not go would throw an error
	*/
	localStr := "localVar"

	fmt.Println(
		"-- non-declared, just initialized variables --",
		"\nnum: ", num, "type: ", reflect.TypeOf(num),
		"\nflt: ", flt, "type: ", reflect.TypeOf(flt),
		"\nstr: ", str, "type: ", reflect.TypeOf(str),

		"\n\n-- just declared variables, with types infered by go --",
		"\ninfNum: ", infNum, "type: ", reflect.TypeOf(infNum),
		"\ninfFlt: ", infFlt, "type: ", reflect.TypeOf(infFlt),
		"\ninfStr: ", infStr, "type: ", reflect.TypeOf(infStr),

		"\n\n-- variable of type type casted to a variable of type int --",
		"\ncastedNum: ", castedNum, "type: ", reflect.TypeOf(castedNum),

		"\n\n-- local variable --",
		"\nlocalStr: ", localStr, "type: ", reflect.TypeOf(localStr),

		"\n\n-- constant variable --",
		"\nconstStr: ", constStr, reflect.TypeOf(constStr),
		"\nconstNum: ", constNum, reflect.TypeOf(constNum),
		"\nconstFlt: ", constFlt, reflect.TypeOf(constFlt),
	)
}
