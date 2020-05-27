package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sli := make([]int, 3, 20)

	cntr := 0
	for {
		fmt.Println("enter integer")
		var x string
		fmt.Scan(&x)
		if x == "X" {
			break
		} else {
			f, err := strconv.Atoi(x)
			if err != nil {
				break
			} else {
				cntr += 1
				if cntr < 4 {
					sli[cntr-1] = f
				} else {
					sli = append(sli, f)
					sort.Ints(sli)
				}
			}

		}
	}
	fmt.Println(sli)

}
