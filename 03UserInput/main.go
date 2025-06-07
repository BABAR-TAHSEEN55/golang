package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// WelcomeMessage := "Welcome to Go Programming"
	// fmt.Println(WelcomeMessage)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Pizza rating :")
	// Comma ok || Err err Syntax
	input, _ := reader.ReadString('\n') // Why can't i use "\n"
	// input, err := reader.ReadString('\n')
	//NOTE : Can give err or _ which means nothing (Also known as Blank Identifier)
	//FIX : Cuz it expects a byte and not a string . Golang doesn't have char
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of Rating :%T", input)

}
