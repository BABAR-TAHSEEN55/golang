package main

import "testing"

func TestIterator(t *testing.T) {
	repeat := Repeat("a", 4)
	expected := "aaaa"
	if repeat != expected {
		t.Errorf("Repeat %q and expected %q", repeat, expected)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a",4)
	}
}
