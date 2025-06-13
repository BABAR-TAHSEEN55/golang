package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Case in Golang")
	fmt.Println("Enter a number between 1-3")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	ConvInput, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	switch ConvInput {
	case 1:
		fmt.Println("It's One")
	case 2:
		fmt.Println("It's two")
	case 3:
		fmt.Println("It's three")
	default:
		fmt.Println("Out of Bound Number")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its Weekday")
	}

}

//TODO : Interface Example
