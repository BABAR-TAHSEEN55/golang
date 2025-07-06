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
func TestArea(t *testing.T) {
	//Helper Functions
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Got %.2f and want %.2f", got, want)
		}

	}

	t.Run("Circles areas ", func(t *testing.T) {
		Circles := Circle{10}
		checkArea(t, Circles, 314.1592653589793)
	})
}
