package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64
	fmt.Print("enter floating pt no")
	fmt.Scan(&x)
	fmt.Println("floating pt no is", x)
	if x > 0 {
		fmt.Println("truncated value is", math.Floor(x))
	} else {
		fmt.Println("truncated value is", math.Ceil(x))

	}

}
