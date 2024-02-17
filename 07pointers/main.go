package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)	// Value of pointer is <nil>

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of pointer is", ptr)	// Value of pointer is 0xc000012100
	fmt.Println("Value of pointer is", *ptr)	// Value of pointer is 23

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber)	// New value is: 25
}
