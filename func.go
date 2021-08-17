package main

import "fmt"

func print(str string) {
	fmt.Println(str)
}

func withParam(name []string, f func(string)) bool {
	if (len(name) > 0) {
		f(name[0])
		return true
	} else {
		fmt.Println("abc")
		return false
	}
}

func returnTwo() (int, bool) {
	return 1, true
}


func demoFunc() {
	withParam([]string {"luigi"}, print)
	withParam([]string {}, print)
	first, second := returnTwo()
	fmt.Println(first, second)
}  