package main

import (
	"fmt"
)

func main() {
	b := true
	for b {
		fmt.Println("Hi")
		b = false
	}

	for a := true; a; a = false {
		fmt.Println(a)
	}

	for key, value := range "suresh" {
		fmt.Println(key, value)
	}

	for _, value := range "suresh" {
		fmt.Println(string(value))
	}
}
