package main

import "fmt"

func main() {
	//defer is used to delay the execution of a function until the surrounding function returns
	//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	//The deferred functions are executed in last-in-first-out order after the surrounding function returns.
	//defer is used to close the file after the function is executed
	defer closeFile()
	openFile()
}

func openFile() {
	fmt.Println("open file")
}

func closeFile() {
	fmt.Println("close file")
}
