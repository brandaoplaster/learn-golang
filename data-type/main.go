package main

import (
	"errors"
	"fmt"
)

func main() {
	// variables int - int8 - int16 - int32 - int64
	var number1 int8 = 100
	fmt.Println(number1)

	number2 := 32
	fmt.Println(number2)

	// variables float - float32 - float64
	var pi float32 = 3.14
	fmt.Println(pi)

	// variables char - return int
	char1 := 'A'
	fmt.Println(char1)

	// variables bool - true - false
	var isTrue bool = true
	fmt.Println(isTrue)
	isFalse := false
	fmt.Println(isFalse)

	// variables error
	var error error
	fmt.Println(error)
	erro := errors.New("Error message")
	fmt.Println(erro)
}
