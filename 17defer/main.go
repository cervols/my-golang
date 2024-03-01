package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// defer fmt.Println("World")
	// // Hello
	// // World

	// defer fmt.Println("World")
	// fmt.Println("Hello")
	// // Hello
	// // World

	// defer works like LIFO
	// defer fmt.Println("World")
	// defer fmt.Println("One")
	// defer fmt.Println("Two")
	// fmt.Println("Hello")
	// // Hello
	// // Two
	// // One
	// // World

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
	// Hello
	// 43210Two
	// One
	// World
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
