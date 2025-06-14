package main

import "fmt"

func main() {
	fmt.Println("Welcome to learning Functions in Golang")
	result := adder(4, 5)
	fmt.Println("Result : ", result)
	ProResult := ProAdder(23, 3, 5, 56)
	fmt.Println("Result : ", ProResult)
}

func adder(ValueOne int, ValueTwo int) int {
	return ValueOne + ValueTwo
}

func ProAdder(Values ...int) int {

	total := 0
	//NOTE: : THis is just adding indexes
	// for i := range Values {
	for _, val := range Values {
		total += val
	}
	return total
}
