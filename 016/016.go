package main

import (
	"fmt"
)

// this program is for lesson 16 in Todd McLeod's Udemy Golang course
// declaring variables & assigning values

func main() {
	a := 42 // the answer to everything
	fmt.Println("I'm printing the 'a' variable below: ")
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, "foo", "bar"
	// note '_' means "foo" will not print
	// _ is a blank identifier so we can track but not use a variable
	// essentially we just throw that return away
	fmt.Println("Printing the values we assigned above: ")
	fmt.Println(b, c, d, f)

	typeExample()
}

func typeExample() {
	/*
		print out a few different variables, int, string, boolean
	*/
	var z int = 1337
	var y string = "31337"
	var w bool = false

	fmt.Println("Printing all 3 values in one statement: ")
	fmt.Println(z, y, w)

	fmt.Printf("\n Printing the values for z,y,w below, newline separated: \n z: %d \n y: %v \n w: %v \n", z, y, w)
}
