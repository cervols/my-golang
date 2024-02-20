package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)	// List of all languages: map[JS:Javascript PY:Python RB:Ruby]
	fmt.Println("JS shorts for:", languages["JS"])	// JS shorts for: Javascript

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)	// List of all languages: map[JS:Javascript PY:Python]


	// loops in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
		// For key JS, value is Javascript
		// For key PY, value is Python
	}

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
	// For key v, value is Javascript
	// For key v, value is Python
}
