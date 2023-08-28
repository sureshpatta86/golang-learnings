package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"Person_Name"`
	Age  int    `json:"Person_Age"`
}

func main() {
	//JSON Marshal and Unmarshal
	suresh := Person{"Suresh", 1}
	fmt.Println(suresh)

	//Marshal returns a byte slice of the JSON encoding of the value passed
	sureshJson, err := json.Marshal(suresh)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sureshJson))

	//Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by the pointer
	var sureshObj Person
	err = json.Unmarshal(sureshJson, &sureshObj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sureshObj)

	//Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by the pointer
	var sureshObj2 map[string]interface{}
	err = json.Unmarshal(sureshJson, &sureshObj2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sureshObj2)

	//Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by the pointer
	var sureshObj3 interface{}
	err = json.Unmarshal(sureshJson, &sureshObj3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sureshObj3)

	var suresh1 Person
	jsonString := `{"Name":"Suresh1","Age":2}`
	err = json.Unmarshal([]byte(jsonString), &suresh1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(suresh1)

	//Different names in struct and json
	var suresh2 Person
	jsonString = `{"Person_Name":"Suresh2","Person_Age":3}`
	err = json.Unmarshal([]byte(jsonString), &suresh2)
	if err != nil {
		fmt.Println(err)
	}
	//we will get zero values for the fields that are not present in the json
	fmt.Println(suresh2) //{ 0}

	//io.Reader and io.Writer
	//json.NewDecoder(r io.Reader) *Decoder
	//json.NewEncoder(w io.Writer) *Encoder

	var aarush Person
	err = json.NewDecoder(os.Stdin).Decode(&aarush)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aarush)

	//Encode
	err = json.NewEncoder(os.Stdout).Encode(aarush)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aarush)

}
