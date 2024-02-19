package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)	// Type of fruitList is []string

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)	// [Apple Tomato Peach Mango Banana]

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)	// [Tomato Peach Mango Banana]

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)	// [Peach Mango]

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777
	// fmt.Println(highScores)	// panic: runtime error: index out of range [4] with length 4

	fmt.Println(highScores)	// [234, 945, 465, 867]

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)	// [234 945 465 867 555 666 321]

	sort.Ints(highScores)
	fmt.Println(highScores)	// [234 321 465 555 666 867 945]
	fmt.Println(sort.IntsAreSorted(highScores))	// true


	// remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)	// [reactjs javascript swift python ruby]
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)	// [reactjs javascript python ruby]
}
