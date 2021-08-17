package main

import "fmt"

func demoLoop () {

	// for loop
	for i:=0; i<5; i++  {
		fmt.Println(i) // 0 1 2 3 4
	}

	names := []string {"mario", "luigi", "peach", "lava"}

	// for loop in array
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for in loop in array
	for index, value := range names {
		fmt.Println(index, value)
	}

	// for in loop in array without using index
	for _, value := range names {
		fmt.Println(value)
	}

	// for in loop in array without using index
	for _, value := range names {
		// "value" below is just a local copy
		value = "abc"
		fmt.Println(value)
	}
	
	// still remain same 
	fmt.Println(names)

	for index, value := range names {
		if index == 0 {
			fmt.Println(index, value)
		} else if index == 1 {
			fmt.Println(index, value)
			continue
		} else {
			fmt.Println(index, value)
			break
		}
		fmt.Println(index, "arrived")
	}

	// switch 
	opt := 1
	switch opt {
	case 1: 
	//
	case 2:
	//
	default:
	//
	}

}