package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()

	// // this is not allowed
	// func greeterTwo() {
	// 	fmt.Println("Another method")
	// }
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is:", result)	// Result is: 8

	// proRes, _ := proAdder(2, 5, 8, 7)
	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is:", proRes)	// Pro result is: 22
	fmt.Println("Pro Message is:", myMessage)	// Pro Message is: Hi Pro result function
}

// function signatures
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values{
		total += val
	}

	return total, "Hi Pro result function"
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Welcome")
}
