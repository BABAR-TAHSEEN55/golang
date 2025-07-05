package main

import "fmt"

// NOTE : Declaring multiple constants together
const (
	EnglishHelloPrefix = "Hello, "
	SpanishHelloPrefix = "Elodie "
	FrenchHelloPrefix  = "Bonjour "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "Super"
	}

	if language == "Spanish" {
		return SpanishHelloPrefix + language
	}
	if language == "French" {
		return FrenchHelloPrefix + language
	}
	return EnglishHelloPrefix + name

}

func main() {
	fmt.Println(Hello("World", "Spanish"))

}

//TODO : Last Refactor
