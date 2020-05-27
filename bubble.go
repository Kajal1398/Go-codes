package main

import "fmt"

func bubble(sli []int) {

	for i := 0; i < 9; i++ {

		for j := 0; j < 10-i-1; j++ {
			if sli[j] >= sli[j+1] {
				Swap(sli, j)

			}

		}
	}

}
func Swap(sli []int, i int) {
	tmp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tmp
}

func main() {

	fmt.Println("enter 10 integers")

	sli := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&sli[i])
	}

	bubble(sli)
	fmt.Println("sorted slice is")
	fmt.Println(sli)
}
