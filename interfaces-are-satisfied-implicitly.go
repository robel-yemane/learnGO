package main

import (
	"fmt"
	"strings"
)

type I interface {
	U() string
	L() string
}

type Str string

//This method means type T implementes the interface I, but we don't need to explicitly declare that it does so.

func (s Str) U() string {
	return strings.ToUpper(string(s))
}

func main() {
	// var i I = T{"robel"}
	// i.M()

	mystr := Str("Forty-Two")
	fmt.Println(mystr)
	up := mystr.U()
	fmt.Println(up)
}
