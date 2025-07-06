package main

import "testing"

//NOTE : Structs are created using {} and not ()

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{3.0, 5.0}
	got := CalculatePerimeter(rectangle)
	want := 16.0
	if got != want {
		t.Errorf("Got %.2f and want %.2f", got, want)
	}

}
func TestArea(t *testing.T) {
	t.Run("Rectangle areas ", func(t *testing.T) {
		rectangle := Rectangle{3.0, 5.0}
		got := rectangle.Area()
		want := 15.0
		if got != want {
			t.Errorf("Got %.2f and want %.2f", got, want)
		}
	})

	t.Run("Circles areas ", func(t *testing.T) {
		Circles := Circle{10}
		got := Circles.Area()
		want := 15.0
		if got != want {
			t.Errorf("Got %g and want %g", got, want)
		}
	})
}
