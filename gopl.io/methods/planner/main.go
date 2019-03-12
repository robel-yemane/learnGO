package main

import (
	"fmt"

	"learnGO/gopl.io/methods/geometry"
)

func main() {
	a := geometry.Point{X: 1, Y: 2}
	b := geometry.Point{X: 4, Y: 5}

	fmt.Println("Distance between two points = ", a.Distance(b))

}
