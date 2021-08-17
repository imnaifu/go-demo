package main

import "fmt"

func demoVariable() {
	// string type variable
	var nameOne string = "mario"

	// inferred string type variable
	var nameTwo = "luigi"

	// declare only and no assignment
	var nameThree string

	// another way to declare and assign
	// can only be used inside function
	nameFour := "peach"

	fmt.Println("string:", nameOne, nameTwo, nameThree, nameFour)

	// change value of a variable
	nameOne = "daisy" 
	fmt.Println("changed value:", nameOne) 

	// ints
	var ageOne int = -20
	var ageTwo uint = 20
	var ageThree = 30
	ageFour := 40
	fmt.Println("int and uint:", ageOne, ageTwo, ageThree, ageFour)
 
	// int with bits
	var numOne int8 = 127
	var numTwo int16 = 2^16
	var numThree uint8 = 255
	fmt.Println("int with bits:",numOne, numTwo, numThree)


	// float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 25.9812312312312312312
	fmt.Println("float:", scoreOne, scoreTwo)

}