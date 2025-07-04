package main

import "fmt"

func GreetGenerator() func() int {
	return func() int {
		fmt.Println("Hellow")
		return 1
	}
}
func main() {
	greet := GreetGenerator()
	fmt.Printf("Greet :%d", greet())

}
