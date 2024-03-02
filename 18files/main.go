package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is:", length)	// length is: 26
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is\n", databyte)
	// Text data inside the file is
	//  [84 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 32 97 32 102 105 108 101]

	fmt.Println("Text data inside the file is\n", string(databyte))
	// Text data inside the file is
	//  This needs to go in a file
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
