package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {

	p := Person{name: "kk", addr: "xncj", phone: "662320"}
	fmt.Print(p)

	// generating json representation from an object
	barr, err := json.Marshal(p) //barr is byte array
	fmt.Print(barr)
	fmt.Print(err)

	//json unmarshaling convts json []byte to a go obj
	var p2 Person
	err1 := json.Unmarshal(barr, &p2)
	fmt.Print(p2)
	fmt.Print(err1)

}
