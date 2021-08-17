package main

import "fmt"

type shape interface {
	area() float64
	circumference() float64
}

// square 
type square struct {
	radius float64
}
func (s square) area() float64 {
	return s.radius
}
func (s square) circumference() float64 {
	return s.radius
}

// circle
type circle struct {
	length float64
}
func (c circle) area() float64 {
	return c.length
}
func (c circle) circumference() float64 {
	return c.length
}

//
func printShape(s shape) {
	fmt.Println("area", s.area())
	fmt.Println("circumference", s.circumference())
}

func demoInterface() {
	shapes := []shape {
		square{radius: 2.0},
		circle{length: 3.0},
		circle{length: 4.0},
		square{radius: 5.0},
	}

	for _, value := range shapes {
		printShape(value)
	}
}