package main

import (
	"fmt"
)

// Variables
// auto dective type
var name = "foo"
var firstName string = "foo"
var lastName string

// global variables
var Name string

// local variables
// constants
const somename = "foo"

// variables declarations

func main() {
	fmt.Println(name)
	fmt.Println(firstName)
	lastName = "foo"
	fmt.Println(lastName)

	// infer the type
	version := 1
	fmt.Println(version)
	fmt.Println(somename)
	localVarb()
}

func localVarb() {
	var so = "foo"
	fmt.Println(so)
}
