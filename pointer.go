package main

import "fmt"

func updateName(n string) {
	n = "mario"
}

func updateNameByPointer(n *string) {
	*n = "mario"
}

func updateMenu(m map[string] float64) {
	m["pie"] = 9.99
}

func demoPointer() {
	// Group A, can be taken as Primitive Value in JS
	// pass by value
	// includes: String, Int, Float, Boolean, Array

	name := "luigi"
	// create a copy of the value of "name" and pass to func
	updateName(name)
	// only copy changed, the original value remains
	fmt.Println(name) // luigi


	// Group B, can be taken as Reference Value in JS
	// pass by pointer
	// includes: Slice, Map, Function
	menu := map[string] float64 {
		"pie": 2.33,
		"coffee": 3.99,
	}
	// crate a copy of the pointer of "menu" and pass to func
	updateMenu(menu)
	// since the copy also point to the "menu", the original value changed
	fmt.Println(menu) // map[coffee:3.99 pie:9.99]

	// get memory address of a variable
	nameAddr := &name
	fmt.Println("memory address of name is:", nameAddr)

	// get value from memory address
	nameValue := *nameAddr
	fmt.Println("value of name is:", nameValue)

	// update by pointer
	// pass the address of "name"
	// in fact, this is what has been done for Slice, Map, Function
	updateNameByPointer(&name)
	fmt.Println("after update by pointer:", name)
}

