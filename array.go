package main

import "fmt"

func demoArray() {
	// array
	var ages [3]int = [3]int {1, 2, 3}
	names := [3]string {"mario", "luigi", "bowser"}
	names[1] = "bo"

	fmt.Println(ages, len(ages)) // [1 2 3] 3
	fmt.Println(names, len(names)) // [mario bo bowser] 3

	// slice (use array under the hood)
	var scores = []int {100, 50, 60} 
	scores[1] = 25
	var newScores = append(scores, 85)

	fmt.Println(scores, len(scores)) // [100 25 60] 3
	fmt.Println(newScores, len(newScores)) // [100 25 60 85] 4

	// slice range
	rangeOne := names[1:2]
	rangeTwo := names[1:]
	rangeThree := names[:2]
	fmt.Println(rangeOne, len(rangeOne)) // [bo] 1
	fmt.Println(rangeTwo, len(rangeTwo)) // [bo bowser] 2
	fmt.Println(rangeThree, len(rangeThree)) // [mario bo] 2

}  