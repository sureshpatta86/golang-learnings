package main

import (
	"fmt"
)

func main() {

	type celsius int
	type fahrenheit int

	// fahrenheit to clesius
	var f fahrenheit = 100

	var c celsius = celsius((f - 32) * 5 / 9) //Formula for converting fahrenheit to celsius

	fmt.Println(f, c)

}
