package main

import(
	"fmt"
	// "math/rand"
	// "time"
	"crypto/rand"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5

	// fmt.Println("The sum is: ", mynumberOne + int(mynumberTwo))	// The sum is:  6

	//random number with math/rand

	// rand.Seed(time.Now().UnixNano())	// after go 1.20 it's not required
	// fmt.Println(rand.Intn(5) + 1)

	//random number from crypto/rand

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
