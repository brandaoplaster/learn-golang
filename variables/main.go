package main

import "fmt"

func main() {
	// variables strings
	var message string = "Variable string"
	varviable1 := "Variable string 1"
	fmt.Println(message)
	fmt.Println(varviable1)

	var (
		name     string = "John"
		lastName string = "Doe"
	)

	fmt.Println(name, lastName)

	color1, color2 := "red", "blue"
	fmt.Println(color1, color2)
}
