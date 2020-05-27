package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ""
	fmt.Print("enter a string")
	fmt.Scan(&str)
	if strings.HasPrefix(str, "i") {
		if strings.Contains(str, "a") {

			var x int = len(str)
			if str[x-1] == 'n' {
				fmt.Print("found")

			} else {
				fmt.Print("not found")
			}

		}
	} else {
		fmt.Print("not found")
	}

}
