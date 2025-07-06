package main

import "testing"

func TestSum(t *testing.T) {
	// t.Run("Coljection of 5 numbers", func(t *testing.T) {
	//
	// 	number := [5]int{1, 2, 3, 4, 5}
	// 	//NOTE : For declaring this type >>
	// 	// number2 := [...]int{1, 2, 3, 4, 5}
	// 	got := Sum(number)
	// 	want := 15
	// 	if got != want {
	// 		//NOTE  : For print -> Default specifier for Arrays : %v
	// 		t.Errorf("Got %d and want %d numbers %v", got, want, number)
	// 	}
	// })
	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			//NOTE  : For print -> Default specifier for Arrays : %v
			t.Errorf("Got %d and want %d numbers %v", got, want, numbers)
		}
	})

}
