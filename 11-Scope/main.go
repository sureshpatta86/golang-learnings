package main

import (
	"fmt"

	"./eu"
)

func main() {
	//Exported variable
	fmt.Println(eu.I)
	//UnExported Variable
	//fmt.Println(eu.i)

	//Block scope
	{
		i := 10
		fmt.Println(i)
	}
	i := 20
	fmt.Println(i)
}
