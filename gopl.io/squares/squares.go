package main

import "fmt"

func squares() func() int {
	var x int
	fmt.Println("I'm here and have x =", x)
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
