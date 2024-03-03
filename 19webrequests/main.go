package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("Handling Web requests in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	// Response is of type: *http.Response

	defer response.Body.Close()	// caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println("Response is:", content)	// html code of the page
}
