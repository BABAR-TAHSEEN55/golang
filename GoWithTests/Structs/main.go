package main

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	length float64
}
type Triangle struct {
	height float64
	length float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length

}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (t Triangle) Area() float64{
	return 0.5 * t.height * t.length
}
func main() {

}
