package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Welcome to time Management with the Gopher")

	presentTime := time.Now()
	// For Formatting , this is the Syntax
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
	createDate := time.Date(2025, time.April, 10, 23, 23, 23, 23, time.UTC)
	fmt.Println(createDate)

	CoreCount := runtime.NumCPU()
	fmt.Println(CoreCount)

}
