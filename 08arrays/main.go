package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is:", fruitList)	// Fruit list is: [Apple Tomato  Peach]
	fmt.Println("Length of fruit list is:", len(fruitList))	// Length of fruit list is: 4

	var vegList = [5]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegy list is:", vegList)	// Vegy list is: [potato beans mushrooms  ]
	fmt.Println("Length of vegy list is:", len(vegList))	// Length of vegy list is: 5
}
