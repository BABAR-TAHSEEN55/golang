package main

import (
	"reflect"
	"testing"
)

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
	t.Run("Slices additon of numbers", func(t *testing.T) {
		got := SumAll([]int{1, 3}, []int{3, 4})
		want := []int{4, 7}
		//NOTE : Can't use != with slices
		// if got != want {
		if reflect.DeepEqual(got, want) {
			t.Errorf("Got %v and want %v", got, want)

		}
	})
}
func TestSumAllTails(t *testing.T) {

	t.Run("Counting tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 3}, []int{3, 4})
		want := []int{3, 4}
		//NOTE : Can't use != with slices
		// if got != want {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v and want %v", got, want)

		}
	})
	t.Run("Counting tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		// if got != want {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v and want %v", got, want)

		}
	})

}
