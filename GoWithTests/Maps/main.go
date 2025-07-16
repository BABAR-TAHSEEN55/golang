package main

import "errors"

var (
	ErrorNotFound          = errors.New("Could not find the meaning of this word")
	ErrorWordExists        = errors.New("Word already exists")
	ErrorWordDoesNotExists = errors.New("Word doesn't exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return value, nil
}
func (d Dictionary) Add(word string, defination string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = defination
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}
func (d Dictionary) Update(word, newDefination string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExists
	case nil:
		d[word] = newDefination
	default:
		return err

	}
	return nil
}
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

// NOTE : Return type as error means i'll return an error if something goes wrong and nil if everything goes right both belongs to error interface
func main() {

}

//TODO : Create a Bank Stimulation -> Complete End to End along with Tests
