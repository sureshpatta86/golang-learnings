package main

import (
	"errors"
	"fmt"
)

// Error Handling
func validate(s string) bool { // return true if valid
	if len(s) == 0 {
		return false
	} else if s == " " {
		return false
	}
	return true
}

func validateWithError(s string) error { //error is a built in interface in go lang which has a method Error() string
	if len(s) == 0 {
		return errors.New("Empty String")
	} else if s == " " {
		return errors.New("String with space")
	}
	return nil
}

func validateWithCustomError(s string) error { //its custom error function
	if len(s) == 0 {
		return customError("Empty String")
	} else if s == " " {
		return customError("String with space")
	}
	return nil
}

// My Custom Error
type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return "My Custom Error:-" + e.Msg
}

func customError(message string) error {
	return &MyError{message}
}

func main() {
	isValid := validate("")
	fmt.Println(isValid)
	if !isValid {
		fmt.Println("Not Valid")
	}

	isValid = validate(" ")
	fmt.Println(isValid)
	if !isValid {
		fmt.Println("Not Valid")
	}

	fmt.Println(validateWithError(""))

	fmt.Println(validateWithError(" "))

	fmt.Println(validateWithCustomError(""))

	fmt.Println(validateWithCustomError(" "))

}
