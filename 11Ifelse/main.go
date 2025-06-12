package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else in Golang ")
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	} else {

		fmt.Println("8 is not divisible by 4")
	}

	if num := 8; num > 9 {
		fmt.Println("Number is greater than 9 ")
	} else if num < 9 {

		fmt.Println("Number is less than 9 ")
	} else {

		fmt.Println("Number is equal than 9 ")
	}

}
