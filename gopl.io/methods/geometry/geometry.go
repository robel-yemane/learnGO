package geometry

import "math"

type Point struct {
	X, Y float64
}

//traditional functions
func Distance(p, q Point) int {
	return int(math.Hypot(q.X-p.X, q.Y-p.Y))
}

//same thing, but as a method of the Point type
func (p Point) Distance(q Point) int {
	return int(math.Hypot(q.X-p.X, q.Y-p.Y))
}
