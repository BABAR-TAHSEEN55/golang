package main

import (
	"fmt"
	"time"
)

func main() {

	//NOTE : Thiese are Anojymous Functions

	// DoStuff()
	// DoStuff = func() {
	// 	fmt.Println("Do Stuff")
	// }
	// DoStuff()
	// DoStuff = func() {
	// 	fmt.Println("Doing some stuff")
	// }
	// DoStuff()
	// Result := InitSeq()
	// fmt.Println(Result())
	// fmt.Println(Result())
	// fmt.Println(Result())
	LimitRequest := RateLimter(5, 5*time.Second)

	for i := 0; i < 8; i++ {
		if LimitRequest() {
			fmt.Printf("Request is Allowed %d\n", i)
		} else {
			fmt.Printf("Request isn't Allowed %d\n", i)
		}
	}
}

var DoStuff func() = func() {
	// DO Stuff
}

func InitSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//NOTE : Real World example of Closures

func RateLimter(maxRequests int, window time.Duration) func() bool {

	requests := 0
	StartTime := time.Now()
	return func() bool {
		if time.Since(StartTime) > window {
			StartTime = time.Now()
			requests = 0
		}
		if requests < maxRequests {
			requests++
			return true
		}
		return false
	}
}
