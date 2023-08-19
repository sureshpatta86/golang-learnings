package main

import (
	"fmt"
)

func main() {
	//If Case
	if s := "suresh"; len(s) > 0 {
		fmt.Println("Lenght of the value is ", len(s))
	}

	//If-Else Case
	a := true
	b := ""
	if a && len(b) > 0 {
		fmt.Println("Both conditions returns true")
	} else {
		fmt.Println("One or more conditions return false")
	}

	//If-ElseIf-Else
	if i := 0; i > 0 {
		fmt.Println("Greater than zero")
	} else if i < 0 {
		fmt.Println("Less than zero")
	} else {
		fmt.Println("Value is zero")
	}
}
