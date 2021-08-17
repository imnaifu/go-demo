package main

import (
	"fmt"
	"sort"
	"strings"
)

func demoLib() {
	greeting := "hello world"

	// get index
	fmt.Println(strings.Index(greeting, "l")) // 2

	// check contains 
	fmt.Println(strings.Contains(greeting, "l")) // true

	// to upper
	fmt.Println(strings.ToUpper(greeting)) // HELLO WORLD

	// replace
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // hi world

	// split
	fmt.Println(strings.Split(greeting, " ")) // [hello world]

	// original value not changed
	fmt.Println(greeting) // hello world

	ages := []int {45, 20, 43, 25, 2, 13, 86, 99}

	// sort int first
	sort.Ints(ages)
	fmt.Println(ages) // [2 13 20 25 43 45 86 99]

	// get location of the int exist
	index1 := sort.SearchInts(ages, 86) 
	fmt.Println(index1) // 6


	// get location of the int to be insert (which not exist now)
	index2 := sort.SearchInts(ages, 3) 
	fmt.Println(index2) // 1


	names := []string {"mario", "luigi", "peach"}

	// sort strings
	sort.Strings(names)
	fmt.Println(names) // [luigi mario peach]

	// get location of the string exist
	fmt.Println(sort.SearchStrings(names, "mario") ) // 1

	// get location of the string to be insert (which not exist now)
	fmt.Println(sort.SearchStrings(names, "a") ) // 0
}