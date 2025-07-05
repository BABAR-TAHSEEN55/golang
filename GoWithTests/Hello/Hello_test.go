package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Goku")
	want := "Hello, Goku"
	if got != want {
		t.Errorf("Got %q and want %q", want, got)
	}

}
