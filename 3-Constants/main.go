package main

import (
	"fmt"
)

//const name string = "suresh"

const (
	name      string = "suresh"
	age       int    = 123
	isMatched bool   = true
)

func main() {
	fmt.Println(name, age, isMatched)

	const number int = 123
	fmt.Println(number)

}
