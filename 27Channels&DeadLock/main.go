package main

import (
	"fmt"
	"sync"
)

// NOTE : Go Routines are not aware on how many of them are running
// NOTE : Channels are used to communicate between Go-Routines( Multiple Go Routines )
// NOTE : 1-> channel listener should be one and so on
func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("Channels in Golang")
	myChnl := make(chan int)
	// Pushing Value into channel
	// myChnl <- 5
	// fmt.Println(<-myChnl)
	wg.Add(2) // ??
	//Receiving Value
	go func(ch chan int, wg *sync.WaitGroup) {
		Value, isChannelOpen := <-myChnl
		// fmt.Println(<-myChnl)
		// fmt.Println(<-myChnl)
		fmt.Println(isChannelOpen)
		fmt.Println(Value)
		wg.Done()
	}(myChnl, wg)
	// Writing values
	go func(ch chan int, wg *sync.WaitGroup) {
		myChnl <- 5
		// myChnl <- 53
		//Closing channel
		//NOTE : Without any value it'll give value as 0 but what if my val is 0? In that case you gotta define a new var which is mentioned above
		close(myChnl)

		wg.Done()
	}(myChnl, wg)
	wg.Wait()
}
