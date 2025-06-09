package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of Pointers")
	myNumber := 23
	// This wont work due to pass by value
	myDuplicateNumber := myNumber + 2

	myPointer := &myNumber
	//This works as it is pass by Reference
	*myPointer = *myPointer + 5

	fmt.Println("Original Value : ", myNumber)
	fmt.Println("Memory Address Value: ", myPointer)
	fmt.Println("myDuplicateNumber : ", myDuplicateNumber)
	fmt.Println("Pointer Value : ", *myPointer)

}
