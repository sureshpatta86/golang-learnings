package main

import "fmt"

type person struct {
	name string
	age  int
}

func (person) sayHello() { // (person) is receiver of sayHello method
	fmt.Println("Hello")
}

func (p *person) updateName(newName string) { // (p *person) is receiver of updateName method
	(*p).name = newName
}

func (p person) updateName1(newName string) { // (p person) is receiver of updateName1 method
	p.name = newName
}

func main() {

	employee := person{name: "John", age: 25}
	fmt.Println(employee)

	employee.age = 30
	fmt.Println(employee)

	employee1 := employee // employee1 is a copy of employee (not a reference)
	fmt.Println(employee1)

	employee1.age = 35
	fmt.Println(employee1)

	employee.sayHello()

	//create a pointer to employee
	employeePointer := &employee
	fmt.Println(employeePointer)

	employeePointer.age = 40
	fmt.Println(employeePointer)
	fmt.Println(employee)

	//update name using pointer
	employeePointer.updateName("Smith")
	fmt.Println(employeePointer)

	//update name using value
	employee.updateName("Smith")
	fmt.Println(employee)

	//update name using value
	employee.updateName1("Joe") // this will not update the name because it is a copy of employee
	fmt.Println(employee)

}
