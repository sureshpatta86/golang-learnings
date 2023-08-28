package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("test.txt")
	writer := io.Writer(file)

	n, err := writer.Write([]byte("Hello, World!\n"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Number of bytes written:", n)
	file.Close()

	// Read file line-by-line.
	file, _ = os.Open("test.txt")
	reader := io.Reader(file)
	buffer := make([]byte, 50)
	n, err = reader.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Number of bytes read:", n)
	file.Close()
}
