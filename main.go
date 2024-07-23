package main

import "fmt"

/**
params
-- How to declare variables
**/

func main() {
	// Var declaration
	var name string = "Christian Chiemela"
	name = "Christian Chiemela"
	var names string
	// Don't declare shorthand outside of the func it won't work
	named := "Christian Chiemela"
	fmt.Println(name, names, named)

	// integer
	var numberOne int = 2
	numberTwo := 4
	fmt.Println(numberOne, numberTwo)

	// bit & memory
	var numberThree int8 = 127
	// unit not negative number
	var numberFour uint = 25
	fmt.Println(numberThree, numberFour)
}
