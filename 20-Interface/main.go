package main

import (
	"fmt"
)

type Animal struct {
	name string
}

type Brids struct {
	name string
}

type eat interface {
	eat()
}

func (a *Animal) eat() {
	fmt.Println("Animal eating")
}

func (b *Brids) eat() {
	fmt.Println("Brids eating")
}

func main() {
	a := []Animal{{"cat"}, {"dog"}}
	fmt.Println(a)

	for index, value := range a {
		fmt.Println(index, value)
		fmt.Println(value.name)
		value.eat()
	}

	b := []Brids{{"crow"}, {"sparrow"}}
	fmt.Println(b)

	for index, value := range b {
		fmt.Println(index, value)
		fmt.Println(value.name)
		value.eat()
	}
	fmt.Println("---------------------------")
	//interface
	c := []eat{&Animal{"cat"}, &Brids{"crow"}, &Animal{"dog"}, &Brids{"sparrow"}}
	fmt.Println(c)

	for index, value := range c {
		fmt.Println(index, value)
		fmt.Println(value)
		value.eat()
	}

}
