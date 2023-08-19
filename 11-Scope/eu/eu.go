package eu

import "fmt"

var exportedVariable = 30
var unexportedVariable = 40

func f() {
	fmt.Println(exportedVariable, unexportedVariable)
}
