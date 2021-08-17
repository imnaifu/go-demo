package main

import "fmt"

func demoPrint() {

	intNum := 61
	floatNum := 61.11

	// pure print
	fmt.Print("hello world")

	// "print line" which print with new line at the end
	fmt.Println("hello world")

	// print formatted strings
	fmt.Printf("my score is %v/%v \n", intNum, 100)

	// print type of variable
	fmt.Printf("my score is type %T \n", intNum)

	// print float type
	fmt.Printf("my score is type %f \n", floatNum)

	// print float type with 2 decimals
	fmt.Printf("my score is type %0.2f \n", floatNum)

	// saved formatted strings
	str := fmt.Sprintf("my score is %v \n", intNum)
	fmt.Print("saved string: ", str)

}