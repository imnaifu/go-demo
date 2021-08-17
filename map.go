package main

import "fmt"

// a map which key is string and value is float
var menu = map[string] float64 {
	"soup": 4.99,
	"pie": 7.99,
	"salad": 6.99,
	"coffee": 3.55,
}

// a map which key is int and value is string
var phonebook = map[int] string {
	1: "mario",
	2: "luigi",
	123: "peach",
}

func demoMap() {
	fmt.Println("menu:", menu)
	 
	// get value by key
	fmt.Println("menu[pie]:", menu["pie"])

	// looping map
	for key, value := range menu {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}

	// modify value
	phonebook[1] = "blabla"
	fmt.Println(phonebook)
}