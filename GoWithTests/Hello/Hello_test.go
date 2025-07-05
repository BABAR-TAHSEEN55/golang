package main

import "testing"

func TestHello(t *testing.T) {
	// t.Run("Saying hello to People", func(t *testing.T) {
	//
	// 	got := Hello("Goku")
	// 	want := "Hello, Goku"
	// 	assertCorrectMessage(t, got, want)
	// })
	// t.Run("Saying hello to nobody", func(t *testing.T) {
	// 	got := Hello("")
	// 	want := Hello("Legend")
	// 	assertCorrectMessage(t, got, want)
	//
	// })
	t.Run("in Spanish ", func(t *testing.T) {
		got := Hello("Elodie ", "Spanish")
		want := "Elodie Spanish"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French ", func(t *testing.T) {
		got := Hello("Bonjour ", "French")
		want := "Bonjour French"
		assertCorrectMessage(t, got, want)
	})

}
func assertCorrectMessage(t testing.TB, got, want string) {
	//NOTE : If you same type like want string, got string ,-> You can give want,got string
	//NOTE : This is a helper func that is satified by both *T and *B
	t.Helper()
	if got != want {
		t.Errorf("Got %q and want %q", want, got)
	}

}

/*
//NOTE :
1) Write a test
2) Make the compiler pass
3) Run the test and see if it fails and check the error message
4) Write enough test to make it pass
5) Refactor

*/
