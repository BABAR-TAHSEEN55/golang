package main

import "testing"

func TestMaps(t *testing.T) {
	dictionary := Dictionary{"name": "Hola my Friend"}
	t.Run("Known Word", func(t *testing.T) {
		got, err := dictionary.Search("name")
		if err != nil {
			t.Fatal("New Error Found")
		}
		want := "Hola my Friend"
		assertStrings(t, got, want)
	})
	t.Run("UnKnown Word", func(t *testing.T) {
		_, got := dictionary.Search("Unknown")
		if got == nil {
			t.Fatal("Error in Unknown error")
		}
		assertError(t, got, ErrorNotFound)
	})
	t.Run("Add a word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "This is the added test")
		want := "This is the added test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("Should find added word", err)
		}
		assertStrings(t, got, want)

	})

}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
func assertError(t testing.TB, want, got error) {
	t.Helper()
	if got != want {
		t.Fatal("expected an error")
	}

}
