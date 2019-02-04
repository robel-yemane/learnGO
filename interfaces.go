package main

import (
	"fmt"
)

// type rect struct {
// 	width, height float64
// }

type geometry interface {
	area() float64
	perim() float64
}

func area() float64 {
	return 3 * 2
}
func perim() float64 {
	return 3 + 2
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	// r := rect{width: 3, height: 4}
	r := geometry.perim
	fmt.Println(r)
}
