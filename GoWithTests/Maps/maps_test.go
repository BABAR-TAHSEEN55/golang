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
		_, err := dictionary.Search("Unknown")
		if err == nil {
			t.Fatal("Error in Unknown error")
		}
		assertError(t, err, ErrorNotFound)
	})
	t.Run("Add a word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defination := "This is the added test"
		dictionary.Add(word, defination)
		// assertStrings(t, got, want)
		assertDefination(t, dictionary, word, defination)

	})
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		defination := "This is just a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")
		assertDefination(t, dictionary, word, defination)
		assertError(t, err, ErrorWordExists)

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
func assertDefination(t testing.TB, dictionary Dictionary, word, defination string) {
	t.Helper()

	got, err := dictionary.Search(word)
	// want := "This is the added test"
	if err != nil {
		t.Fatal("Should find added word", err)
	}
	assertStrings(t, got, defination)
	//NOTE : Here defination becomes your want
}
func TestUpdate(t *testing.T) {
	t.Run("Existing Word", func(t *testing.T) {

		word := "test"
		defination := "this is just a test"
		Dictionary := Dictionary{word: defination}
		newDefination := "This is the updated Test"
		Dictionary.Update(word, newDefination)
		assertDefination(t, Dictionary, word, newDefination)
	})
	t.Run("New Word", func(t *testing.T) {
		word := "New Word"
		defination := "This is the new word"
		Dictionary := Dictionary{}
		err := Dictionary.Update(word, defination)
		assertError(t, err, ErrorWordDoesNotExists)
	})

}
func TestDelete(t *testing.T) {
	word := "New Word"
	Dictionary := Dictionary{word: "Test defination"}
	err := Dictionary.Delete(word)
	assertError(t, err, nil)
	_, err = Dictionary.Search(word) // No need to use := as it has been already declared but what about scope?
	assertError(t, err, ErrorNotFound)

}
