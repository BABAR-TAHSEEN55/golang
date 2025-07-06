package main

import "testing"

//NOTE : Structs are created using {} and not ()

// func TestPerimeter(t *testing.T) {
//
//	rectangle := Rectangle{3.0, 5.0}
//	got := CalculatePerimeter(rectangle)
//	want := 16.0
//	if got != want {
//		t.Errorf("Got %.2f and want %.2f", got, want)
//	}
//
// }
// New way of testing without creating individual tests
func TestArea(t *testing.T) {
	//table driven Tests
	areatTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		//Either is fine
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangel", Triangle{2, 4}, 4.0},
	}
	for _, tt := range areatTests {
		//Using tt.name and declaring it instead of manually writing "Rectangel" etc
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("Got %g and want %g", got, tt.want)
			}
		})
	}

}
