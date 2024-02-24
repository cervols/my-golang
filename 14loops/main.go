package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)	// [Sunday Tuesday Wednesday Friday Saturday]

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	// Sunday
	// Tuesday
	// Wednesday
	// Friday
	// Saturday

	for i := range days {
		fmt.Println(days[i])
	}
	// Sunday
	// Tuesday
	// Wednesday
	// Friday
	// Saturday

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}
	// index is 0 and value is Sunday
	// index is 1 and value is Tuesday
	// index is 2 and value is Wednesday
	// index is 3 and value is Friday
	// index is 4 and value is Saturday


	// BREAK
	rogueValue := 1

	for rogueValue < 10 {	// equivalent of WHILE in other languages
		if rogueValue == 5 {
			break
		}

		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}
	// Value is: 1
	// Value is: 2
	// Value is: 3
	// Value is: 4


	// CONTINUE
	rogueValue = 1

	for rogueValue < 10 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}
	// Value is: 1
	// Value is: 2
	// Value is: 3
	// Value is: 4
	// Value is: 6
	// Value is: 7
	// Value is: 8
	// Value is: 9


	// GOTO
	rogueValue = 1

	for rogueValue < 10 {
		if rogueValue == 3 {
			goto lco
		}

		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}
	// Value is: 1
	// Value is: 2
	// Jumping at website.com

	// LABELS
	lco:
		fmt.Println("Jumping at website.com")
}
