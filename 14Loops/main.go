package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang")
	days := []string{"Sunday", "Monday", "Tuesday", "Thrusday"}

	// 1st Way
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// 	fmt.Println("I : ", i)
	// }
	// 2nd way
	for i := range days {
		if i == 2 {
			fmt.Println("Skipped day : ", days[2])
			continue
		}
		fmt.Println(days[i])
	}
	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}
		fmt.Println("Value : ", rogueValue)
		rogueValue++
	}
	// This is known as Label
lco:
	fmt.Println("Jumping to www.google.com")

}
