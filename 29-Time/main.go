package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)

	fmt.Println(t.Day())
	fmt.Println(t.Month())
	fmt.Println(t.Year())

	fmt.Println(t.UTC())

	fmt.Println(t.Weekday() == time.Friday)
	t1 := time.Date(2023, time.August, 01, 10, 10, 10, 10, time.Now().Location())
	fmt.Println(t1)
	t2 := time.Now()
	time.Sleep(time.Second * 2)
	fmt.Println(time.Since(t2))
}
