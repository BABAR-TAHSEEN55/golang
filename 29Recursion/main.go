package main

import "fmt"

func Fib(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fib(n-1)
}
func main() {
	fmt.Println("Fibonacci : ", Fib(7))

}
