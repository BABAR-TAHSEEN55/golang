package main

import "fmt"

func main() {
	defer fmt.Println("I'll be printed at the last ")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hellow ? ")
	fmt.Println("Welcome to Defers")
	MyDefer()

}
func MyDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
