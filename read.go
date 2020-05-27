package main

import (
	"fmt"
	"io/ioutil"
)

type Person struct {
	fname string
	lname string
}

func main() {

	fmt.Println("enter name of file")
	var filenm string
	fmt.Scan(&filenm)
	data, err := ioutil.ReadFile(filenm)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Print(string(data))
}
