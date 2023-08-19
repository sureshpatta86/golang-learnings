package main

import (
	"fmt"
)

func main() {
	//Condition based switching
	i := 0
	switch {
	case i > 0:
		fmt.Println("Value is greater than zero", i)
	case i < 0:
		fmt.Println("Value is less than zero", i)
	default:
		fmt.Println("Value is zero", i)
	}

	//value base switching
	name := ""
	switch name {
	case "suresh", "nani":
		fmt.Println("Its my name", name)
	case "p":
		fmt.Println("Its my name", name)
	case "":
		fmt.Println("Name value is empty")
	default:
		fmt.Println("Its default name", name)
	}
}
