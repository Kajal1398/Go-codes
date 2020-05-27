package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("kk.go")

	fmt.Print(string(data))
	fmt.Print(err)

	dat := []byte("hello world")
	err1 := ioutil.WriteFile("out.txt", dat, 0777)
	fmt.Print(err1)

	f, e := os.Open("out.txt")
	barr := make([]byte, 10)
	nb, ee := f.Read(barr)

	fmt.Println(e, ee, nb)
	f.Close()

	f1, errr := os.Create("of.txt")
	br := []byte{1, 2, 3, 4}
	nb1, e1 := f1.Write(br)
	nb2, e2 := f1.WriteString("Hi")

	fmt.Println(errr, e1, e2, nb1, nb2)
}
