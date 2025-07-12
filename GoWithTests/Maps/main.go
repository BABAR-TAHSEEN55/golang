package main

import "errors"

var ErrorNotFound = errors.New("Could not find the meaning of this word")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return value, nil
}
func (d Dictionary) Add(word string, defination string) {
	d[word] = defination

}
func main() {

}
