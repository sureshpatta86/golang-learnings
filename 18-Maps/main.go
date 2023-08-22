package main

import "fmt"

func main() {
	var m map[string]string //map initiated

	fmt.Println(m, len(m))
	if m == nil {
		fmt.Println("Map default value is nil")
	}

	m = map[string]string{"suresh": "myname", "nani": "mypetname"} //key:value,key:value
	fmt.Println(m, len(m))

	fmt.Println(m["suresh"]) //get the value based on the key
	m["suresh"] = "my-name"  //assign the new value to existing key
	fmt.Println(m["suresh"])

	m["p"] = "sname" // add new key:value to map
	fmt.Println(m, len(m))

	delete(m, "p") //delete the key:value from map
	fmt.Println(m, len(m))

	for key, value := range m { //iterate the map using for loop and range
		fmt.Println(key, value)
	}

	value := m["suresh"] //get the value based on the key
	fmt.Println(value)

	value1 := m["abc"]
	fmt.Println(value1)

	value, ok := m["abc"] //get the value based on the key and check the key is present or not
	fmt.Println(value, ok)

	//pre allocation

	m = make(map[string]string, 5) //pre allocation of map with 5 key:value

	fmt.Println(m, len(m))

}
