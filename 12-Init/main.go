package main

import (
	"fmt"

	"./i"
)

func main() {
	fmt.Println(i.I)
	fmt.Println("Hello World")
}

func init() {
	fmt.Println("Init1")
}

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init3")
}
