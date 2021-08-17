package main

import "fmt"

type bill struct {
	name string
	item map[string] float64
	tip float64
}

// construct function
func billConstructor(name string) bill {
	b := bill {
		name: name,
	}
	return b
}

// receiver functions
// format to string
func (b *bill) format() string {
	// here we pass "b" as a reference 
	// so that we don't need to make a copy every time when calling this func
	// only make a copy of the pointer instead
	formatted := fmt.Sprintf("bill name is %v", b.name)
	return formatted
}

// here need to pass b as pointer, or can not update the "tip"
func (b *bill ) updateTip(tip float64) {
	// dereference b
	// (*b).tip = tip  

	// golang automatic dereference the "b"
	b.tip = tip  
}

func demoStruct() {
	myBill := billConstructor("abc")
	myBill.updateTip(2.33)
	fmt.Println(myBill) // {abc map[] 0}

	myBillFormatted := myBill.format()
	fmt.Println(myBillFormatted) // bill name is abc
}