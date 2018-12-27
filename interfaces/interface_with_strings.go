package main

import (
	"fmt"
)

type Caser interface {
	Mum() string
	Dad() string
}

type myStr string

func (s myStr) Mum() {
	fmt.Println()
}

func main() {

	nateyString := myStr(33)
	nateyString.Mum()
}

// type CustomString string

// const (
// 	Foobar CustomString = "custom string"
// )

// func (c CustomString) String() string {
// 	fmt.Println("Executing String() for CustomString")
// 	return string(c)
// }

// func StringReturn() string {
// 	return Foobar.String()
// }

// func main() {
// 	fmt.Printf("%q\n,%g\n,%v\n,%T\n", Foobar, Foobar, Foobar, Foobar)

// }
