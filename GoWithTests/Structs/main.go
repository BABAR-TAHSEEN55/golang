package main

type Rectangle struct {
	width  float64
	length float64
}
type Circle struct {
	radius float64
}

func CalculatePerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.length + rectangle.width)

}
func CalculateArea(rectangle Rectangle) float64 {
	return rectangle.length * rectangle.width
}
func main() {

}
