package main

import "fmt"

func main() {
	//Defer initianlization
	defer CleanUp()

	//Panic
	const a = 11
	if a != 10 {
		panic("a is not 10")
	}
}

func CleanUp() {
	//Defer
	if r := recover(); r != nil {
		fmt.Println("Defer CleanUp-", r)
	}
}
