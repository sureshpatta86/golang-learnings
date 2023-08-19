package main

import (
	"fmt"

	eu "github.com/suresh-patta/GoLangMaterial/11-Scope/eu"
)

func main() {
	//Exported variable
	fmt.Println(eu.ExportedVariable)
	fmt.Println(eu.unexportedVariable)

	//Block scope
	{
		i := 10
		fmt.Println(i)
	}
	i := 20
	fmt.Println(i)
}
