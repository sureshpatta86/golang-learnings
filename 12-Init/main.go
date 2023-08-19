package main

import (
	"fmt"

	"./oi"
)

func main() {
	fmt.Println(oi.I)
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
