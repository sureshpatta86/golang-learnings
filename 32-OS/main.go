package main

import (
	"fmt"
	"os"
)

func main() {

	n, err := os.Stdout.WriteString("Hello, World!\n")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Number of bytes written:", n)

	dirName := "test"
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fileName := dirName + "/test.txt"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer f.Close()

	// Write some text line-by-line to file.
	l, err := f.WriteString("Hello, World!\n")
	if err != nil {
		fmt.Println("Error:", err)
	}
	f.Close()
	fmt.Println("Number of bytes written:", l)

	// Save file changes.
	err = f.Sync()
	if err != nil {
		fmt.Println("Error:", err)
	}
	f.Close()

	// Read file line-by-line.
	f, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	buffer := make([]byte, 50)
	n, err = f.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Number of bytes read:", n)
	f.Close()

	//Delete file.
	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//Delete directory.
	err = os.Remove(dirName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(os.Getwd())        //get current working directory
	fmt.Println(os.Getenv("PATH")) //get environment variable
	fmt.Println(os.Hostname())     //get hostname
	fmt.Println(os.Environ())      //get all environment variables
}
